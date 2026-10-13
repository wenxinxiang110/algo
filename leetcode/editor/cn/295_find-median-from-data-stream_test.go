package main

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     [][]int
		want     []float64
	}{
		{
			name:     "基础用例 - LeetCode示例1",
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			args:     [][]int{{}, {1}, {2}, {}, {3}, {}},
			want:     []float64{1.5, 2.0},
		},
		{
			name:     "单个元素",
			commands: []string{"MedianFinder", "addNum", "findMedian"},
			args:     [][]int{{}, {5}, {}},
			want:     []float64{5.0},
		},
		{
			name:     "两个相同元素",
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian"},
			args:     [][]int{{}, {3}, {3}, {}},
			want:     []float64{3.0},
		},
		{
			name:     "四个元素",
			commands: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			args:     [][]int{{}, {1}, {2}, {3}, {4}, {}},
			want:     []float64{2.5},
		},
		{
			name:     "负数元素",
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian"},
			args:     [][]int{{}, {-1}, {-2}, {}},
			want:     []float64{-1.5},
		},
		{
			name:     "交替插入",
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			args:     [][]int{{}, {5}, {1}, {}, {3}, {}, {7}, {}},
			want:     []float64{3.0, 3.0, 4.0},
		},
		{
			name:     "降序插入",
			commands: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			args:     [][]int{{}, {5}, {4}, {3}, {}},
			want:     []float64{4}, // 当前实现返回3.0而不是4.0
		},
		{
			name:     "重复数字",
			commands: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			args:     [][]int{{}, {2}, {2}, {2}, {2}, {}},
			want:     []float64{2.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mf *MedianFinder
			wantIndex := 0

			for i, cmd := range tt.commands {
				switch cmd {
				case "MedianFinder":
					obj := Constructor()
					mf = &obj
				case "addNum":
					if mf == nil {
						t.Fatal("MedianFinder not initialized before addNum")
					}
					if len(tt.args[i]) != 1 {
						t.Fatalf("addNum requires exactly 1 argument, got %d", len(tt.args[i]))
					}
					mf.AddNum(tt.args[i][0])
				case "findMedian":
					if mf == nil {
						t.Fatal("MedianFinder not initialized before findMedian")
					}
					got := mf.FindMedian()

					if wantIndex >= len(tt.want) {
						t.Fatalf("unexpected findMedian call, no more expected values")
					}

					want := tt.want[wantIndex]
					// 允许浮点数比较的误差范围
					if abs(got-want) > 1e-5 {
						t.Errorf("findMedian() = %v, want %v", got, want)
					}
					wantIndex++
				default:
					t.Fatalf("unknown command: %s", cmd)
				}
			}

			// 验证所有期望的findMedian调用都已完成
			if wantIndex != len(tt.want) {
				t.Errorf("expected %d findMedian calls, but only executed %d", len(tt.want), wantIndex)
			}
		})
	}
}

// 辅助函数：计算浮点数的绝对值
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// 性能测试：测试大量操作的性能
func BenchmarkMedianFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := Constructor()
		// 插入1000个数字
		for j := 0; j < 1000; j++ {
			mf.AddNum(j)
			if j%100 == 0 {
				_ = mf.FindMedian()
			}
		}
	}
}
