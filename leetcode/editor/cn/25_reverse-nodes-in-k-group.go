package main

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 2751 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		//找到尾节点
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 尾节点没循环结束就为空，说明已经遍历完最后一组了
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		//head, tail = reversePart(head, tail)
		head, tail = reverseListPreK(head, k)

		// 续上头和尾巴
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next

	}

	return hair.Next
}

// 反转 head -> tail部分的链表，并且返回新的头和尾
// tail稍微有点多余, 在原有的 reverseList 中需要额外支持返回一个尾节点
func reversePart(head, tail *ListNode) (*ListNode, *ListNode) {
	// 这里其实是一个虚拟头节点，不管初始化多少，都不影响
	prev := /*tail.Next*/ &ListNode{}

	// 翻转，这里是用的尾插法
	iter := head
	for prev != tail {
		nex := iter.Next
		iter.Next = prev
		prev = iter
		iter = nex

		// iter.Next, = prev, iter.Next, prev,iter

	}
	return tail, head
}

// 翻转链表前 k 个元素，返回该组的第一个和最后一个节点
func reverseListPreK(head *ListNode, k int) (newHead, newEnd *ListNode) {
	var pre *ListNode = nil
	newEnd = head
	//var trueNext *ListNode
	for i := 1; i <= k && head != nil; i++ {
		pre, head, head.Next = head, head.Next, pre
		//trueNext = head.Next
		//head.Next = pre
		//pre = head
		//head = trueNext
	}
	newHead = pre
	return
}

//leetcode submit region end(Prohibit modification and deletion)
