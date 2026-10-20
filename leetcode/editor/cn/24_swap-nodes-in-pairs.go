package main

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
// Related Topics 递归 链表 👍 2293 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return swapPairsDirect(head)
}

func swapPairEasy(head *ListNode) *ListNode {
	var iter = &ListNode{Next: head}
	var newHead = iter
	for iter.Next != nil && iter.Next.Next != nil {
		iter, iter.Next, iter.Next.Next, iter.Next.Next.Next =
			iter.Next, iter.Next.Next, iter.Next.Next.Next, iter.Next
	}
	return newHead.Next
}

func swapPairsDirect(head *ListNode) *ListNode {
	var iter = &ListNode{Next: head}
	var newHead = iter
	for iter.Next != nil && iter.Next.Next != nil {
		// 1. 记录 第二个 ele 位置
		tmp := iter.Next.Next
		// 2. 记录 下一组位置
		tmp2 := iter.Next.Next.Next
		// 断开第一个元素的指针，指向下一组
		iter.Next.Next = tmp2
		// 断开第二个的指针，指回第一个
		tmp.Next = iter.Next
		// 上一组(哨兵)指回新头 B，保证整条链不断
		iter.Next = tmp
		// iter进入到下一组（移动到本组交换后的尾节点 A）
		iter = tmp.Next
	}
	return newHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
