package heap

import (
	"math/rand"
	"reflect"
	"testing"
)

// opKind 表示一次堆操作类型。
type opKind int

const (
	pushOp opKind = iota // 向堆中压入 arg 指定的元素
	popOp                // 弹出堆顶，期望返回值等于 want
	topOp                // 查看堆顶，期望返回值等于 want
	lenOp                // 查询堆大小，期望等于 want
)

// op 描述一次对堆的操作及预期结果。
type op struct {
	kind opKind
	arg  int // pushOp 时生效：要压入的元素
	want int // popOp/topOp 时生效：期望返回值；lenOp 时生效：期望长度
}

// dontCheck 作为 want 的哨兵值，表示“执行该操作但不校验返回值”。
// 仅用于与具体堆序无关的行为契约用例（如验证 Top 不弹出、Len 变化）。
const dontCheck = -1 << 31

// newHeapFunc 构造一个 int 堆（满足 Interface[int]），供测试驱动。
// 之后你实现自己的堆时，只需提供一个返回 Interface[int] 的构造函数即可复用本测试。
type newHeapFunc func() Interface[int]

// implKind 标记一种堆实现的语义，用于把用例匹配到正确的实现上。
type implKind int

const (
	kindMin implKind = iota // 最小堆（小值优先）
	kindMax                 // 最大堆（大值优先）
	kindAbs                 // 按绝对值比较（小绝对值优先）
	kindAll                 // 通用：所有实现都跑
)

// impl 描述一种待测实现及其语义。
type impl struct {
	name string
	kind implKind
	make func() Interface[int]
}

// runOps 按顺序执行操作序列并对结果做断言。
func runOps(t *testing.T, h Interface[int], ops []op) {
	t.Helper()
	for i, o := range ops {
		switch o.kind {
		case pushOp:
			h.Push(o.arg)
		case popOp:
			got := h.Pop()
			if o.want != dontCheck && got != o.want {
				t.Fatalf("op#%d Pop() = %d, want %d", i, got, o.want)
			}
		case topOp:
			got := h.Top()
			if o.want != dontCheck && got != o.want {
				t.Fatalf("op#%d Top() = %d, want %d", i, got, o.want)
			}
		case lenOp:
			if got := h.Len(); got != o.want {
				t.Fatalf("op#%d Len() = %d, want %d", i, got, o.want)
			}
		}
	}
}

// TestHeap 表驱动测试：每个用例声明适用的最小堆/最大堆/自定义比较实现，
// 只对匹配语义的实现校验，从而验证各实现的正确堆序。
func TestHeap(t *testing.T) {
	// 注册所有待测实现：你之后自己的实现也加到这里即可（填对应 kind）。
	impls := []impl{
		{name: "MinHeap", kind: kindMin, make: func() Interface[int] { return MinHeap[int]() }},
		{name: "MaxHeap", kind: kindMax, make: func() Interface[int] { return MaxHeap[int]() }},
		{name: "New(abs)", kind: kindAbs, make: func() Interface[int] {
			return New[int](func(a, b int) bool { return abs(a) < abs(b) })
		}},
	}

	tests := []struct {
		name      string
		appliesTo implKind // 仅对语义匹配的实现运行该用例
		ops       []op
	}{
		// ===== 最小堆（kindMin）专属场景 =====
		{
			name:      "最小堆_顺序插入升序弹出",
			appliesTo: kindMin,
			ops: []op{
				{lenOp, 0, 0},
				{pushOp, 5, 0}, {pushOp, 1, 0}, {pushOp, 3, 0}, {lenOp, 0, 3},
				{popOp, 0, 1}, {popOp, 0, 3}, {popOp, 0, 5}, {lenOp, 0, 0},
			},
		},
		{
			name:      "最小堆_逆序插入升序弹出",
			appliesTo: kindMin,
			ops: []op{
				{pushOp, 9, 0}, {pushOp, 7, 0}, {pushOp, 8, 0}, {pushOp, 6, 0},
				{popOp, 0, 6}, {popOp, 0, 7}, {popOp, 0, 8}, {popOp, 0, 9},
			},
		},
		{
			name:      "最小堆_负数参与排序",
			appliesTo: kindMin,
			ops: []op{
				{pushOp, -3, 0}, {pushOp, -1, 0}, {pushOp, -2, 0},
				{popOp, 0, -3}, {popOp, 0, -2}, {popOp, 0, -1},
			},
		},
		// ===== 最大堆（kindMax）专属场景 =====
		{
			name:      "最大堆_弹出应为降序",
			appliesTo: kindMax,
			ops: []op{
				{pushOp, 1, 0}, {pushOp, 5, 0}, {pushOp, 3, 0}, {pushOp, 2, 0},
				{popOp, 0, 5}, {popOp, 0, 3}, {popOp, 0, 2}, {popOp, 0, 1},
			},
		},
		{
			name:      "最大堆_负数参与排序",
			appliesTo: kindMax,
			ops: []op{
				{pushOp, -3, 0}, {pushOp, -1, 0}, {pushOp, -2, 0},
				{popOp, 0, -1}, {popOp, 0, -2}, {popOp, 0, -3},
			},
		},
		// ===== 自定义比较（kindAbs）专属场景 =====
		{
			name:      "自定义比较_按绝对值最小优先",
			appliesTo: kindAbs,
			ops: []op{
				{pushOp, 5, 0}, {pushOp, -1, 0}, {pushOp, 3, 0}, {pushOp, -4, 0},
				{topOp, 0, -1},                // 绝对值 1 最小
				{popOp, 0, -1}, {popOp, 0, 3}, // 绝对值 3
				{popOp, 0, -4}, {popOp, 0, 5},
			},
		},
		// ===== 通用场景：所有实现都应满足（与具体堆序无关，只验证行为契约）=====
		{
			name:      "通用_Top不弹出且Len不变",
			appliesTo: kindAll, // 通用，任意实现取自身语义校验行为契约
			ops: []op{
				{pushOp, 4, 0}, {pushOp, 2, 0}, {pushOp, 6, 0}, {lenOp, 0, 3},
				{topOp, 0, dontCheck}, // 仅校验 Top 不 panic 且 Top 不弹出
				{lenOp, 0, 3},         // Top 后长度仍应为 3（证明 Top 未弹出）
				{popOp, 0, dontCheck}, // 弹出一个（值不校验，堆序由专属用例覆盖）
				{lenOp, 0, 2},         // 弹出后长度变 2
			},
		},
		{
			name:      "通用_重复元素",
			appliesTo: kindAll,
			ops: []op{
				{pushOp, 2, 0}, {pushOp, 2, 0}, {pushOp, 2, 0},
				{popOp, 0, 2}, {popOp, 0, 2}, {popOp, 0, 2}, {lenOp, 0, 0},
			},
		},
		{
			name:      "通用_单元素堆",
			appliesTo: kindAll,
			ops: []op{
				{pushOp, 42, 0}, {lenOp, 0, 1}, {topOp, 0, 42},
				{popOp, 0, 42}, {lenOp, 0, 0},
			},
		},
		{
			name:      "通用_空堆边界_PopTop返回零值",
			appliesTo: kindAll,
			ops: []op{
				{lenOp, 0, 0},
				{popOp, 0, 0}, // 空堆 Pop 返回零值
				{topOp, 0, 0}, // 空堆 Top 返回零值
				{lenOp, 0, 0},
			},
		},
	}

	// 把"通用"用例复制到每种实现都跑，专属用例只跑匹配实现。
	for _, im := range impls {
		t.Run(im.name, func(t *testing.T) {
			for _, tt := range tests {
				// 专属用例：仅当语义匹配时运行；通用用例(kindAll)所有实现都跑
				if tt.appliesTo != im.kind && tt.appliesTo != kindAll {
					continue
				}
				t.Run(tt.name, func(t *testing.T) {
					runOps(t, im.make(), tt.ops)
				})
			}
		})
	}
}

// TestHeap_PopOrder 用"逐步弹出并收集"的方式做端到端验证：
// 最小堆应得到升序序列，最大堆应得到降序序列。
func TestHeap_PopOrder(t *testing.T) {
	impls := []impl{
		{name: "MinHeap", kind: kindMin, make: func() Interface[int] { return MinHeap[int]() }},
		{name: "MaxHeap", kind: kindMax, make: func() Interface[int] { return MaxHeap[int]() }},
	}

	tests := []struct {
		name     string
		wantDesc bool // true=期望降序(最大堆)，false=期望升序(最小堆)
	}{
		{name: "随机100_最小堆升序", wantDesc: false},
		{name: "随机100_最大堆降序", wantDesc: true},
	}

	for _, im := range impls {
		for _, tt := range tests {
			// 最小堆只验证升序用例，最大堆只验证降序用例
			if (im.kind == kindMin) != !tt.wantDesc {
				continue
			}
			t.Run(im.name+"/"+tt.name, func(t *testing.T) {
				input := randSeq(100)
				h := im.make()
				for _, v := range input {
					h.Push(v)
				}
				got := make([]int, 0, h.Len())
				for h.Len() > 0 {
					got = append(got, h.Pop())
				}
				want := append([]int(nil), input...)
				sortInts(want)
				if tt.wantDesc {
					reverse(want)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("pop order = %v, want %v", got, want)
				}
			})
		}
	}
}

// ---- 测试辅助 ----

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func randSeq(n int) []int {
	r := rand.New(rand.NewSource(42))
	s := make([]int, n)
	for i := range s {
		s[i] = r.Intn(1000) - 500 // [-500, 499]
	}
	return s
}

func sortInts(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
