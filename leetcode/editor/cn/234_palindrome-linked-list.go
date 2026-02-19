package main

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 2201 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return isPalindromeBest(head)
}

func isPalindromeBest(head *ListNode) bool {
	//
	if head == nil {
		return true
	}

	// find first half
	var firstEnd = endOfFirstHalf(head)

	// 反转第二部分链路
	var reverseList = func(inline *ListNode) *ListNode {
		var pre *ListNode = nil
		for inline != nil {
			pre, inline, inline.Next = inline, inline.Next, pre
		}
		return pre
	}
	var secondEnd = reverseList(firstEnd.Next)

	//判断是否回文
	for secondEnd != nil {
		if secondEnd.Val != head.Val {
			return false
		}
		secondEnd = secondEnd.Next
		head = head.Next
	}
	return true
}

func endOfFirstHalf(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
