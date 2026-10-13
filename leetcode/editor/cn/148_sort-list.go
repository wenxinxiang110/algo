package main

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 2702 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sortListTopdown(head)
}

// 归并排序: 自顶向下
func sortListTopdown(head *ListNode) *ListNode {
	// 找到 tail
	return sortListTopdownRecursive(head, nil)
}

// 归并排序: 自顶向下 返回头节点;注意
func sortListTopdownRecursive(head, tail *ListNode) *ListNode {
	// 空链表
	if head == nil {
		return head
	}

	// 只有一个节点直接返回，注意，tail 是不包括在排序内的
	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 找到中点: 快慢指针
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 断开链表，避免循环引用
	mid := slow

	// 递归排序
	list1 := sortListTopdownRecursive(head, mid)
	list2 := sortListTopdownRecursive(mid, tail)

	// 合并
	sorted := mergeTwoLists4sortList(list1, list2)
	return sorted
}

func mergeTwoLists4sortList(l1, l2 *ListNode) *ListNode {
	iter := &ListNode{}
	pre := iter
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			iter.Next = l1
			l1 = l1.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
		}
		iter = iter.Next
	}
	remind := l1
	if remind == nil {
		remind = l2
	}
	iter.Next = remind
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
