package main

import (
	"testing"
)

// buildList 由切片构造链表，返回头节点
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

// getTail 返回链表尾节点
func getTail(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

// nodeAt 返回第 index 个节点（0-based），越界返回 nil
func nodeAt(head *ListNode, index int) *ListNode {
	cur := head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	return cur
}

func Test_getIntersectionNode(t *testing.T) {
	tests := []struct {
		name  string
		headA []int
		headB []int
		// 在 A 中（从头起）跳到交点的节点下标；-1 表示不相交
		skipA int
		// 在 B 中（从头起）跳到交点的节点下标；-1 表示不相交
		skipB int
		// 期望返回的交点值；-1 表示期望返回 nil
		wantVal int
	}{
		{
			name:    "相交于中后段",
			headA:   []int{4, 1, 8, 4, 5},
			headB:   []int{5, 6, 1, 8, 4, 5},
			skipA:   2,
			skipB:   3,
			wantVal: 8,
		},
		{
			name:    "相交于中后段(示例2)",
			headA:   []int{1, 9, 1, 2, 4},
			headB:   []int{3, 2, 4},
			skipA:   3,
			skipB:   1,
			wantVal: 2,
		},
		{
			name:    "相交于起点(两链表共享头节点)",
			headA:   []int{1, 2, 3},
			headB:   []int{1, 2, 3},
			skipA:   0,
			skipB:   0,
			wantVal: 1,
		},
		{
			name:    "不相交",
			headA:   []int{2, 6, 4},
			headB:   []int{1, 5},
			skipA:   -1,
			skipB:   -1,
			wantVal: -1,
		},
		{
			name:    "其中一个为nil",
			headA:   []int{1, 2, 3},
			headB:   nil,
			skipA:   -1,
			skipB:   -1,
			wantVal: -1,
		},
		{
			name:    "两个都为nil",
			headA:   nil,
			headB:   nil,
			skipA:   -1,
			skipB:   -1,
			wantVal: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA := buildList(tt.headA)
			headB := buildList(tt.headB)

			var want *ListNode
			if tt.wantVal == -1 {
				want = nil
			} else {
				// 构造真实相交：把 B 的尾部接到 A 的 skipA 处（共享内存）
				if tt.skipA >= 0 {
					intersectA := nodeAt(headA, tt.skipA)
					// B 尾部接到交点，形成真实相交
					if tailB := getTail(headB); tailB != nil {
						tailB.Next = intersectA
					}
					want = intersectA
				} else {
					want = nil
				}
			}

			got := getIntersectionNode(headA, headB)

			// 相交场景：必须返回完全相同的节点指针且值匹配
			if want != nil {
				if got != want {
					t.Errorf("getIntersectionNode() 指针不相等: got=%p(%v), want=%p(%v)", got, gotVal(got), want, tt.wantVal)
				}
				if got != nil && got.Val != tt.wantVal {
					t.Errorf("getIntersectionNode() 值=%v, want=%v", got.Val, tt.wantVal)
				}
				return
			}

			// 不相交场景：必须返回 nil（同时验证不会死循环）
			if got != nil {
				t.Errorf("getIntersectionNode() = %p(%v), want nil", got, gotVal(got))
			}
		})
	}
}

func gotVal(n *ListNode) int {
	if n == nil {
		return -1
	}
	return n.Val
}
