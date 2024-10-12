package main

//给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：能尝试使用一趟扫描实现吗？
//
//
//
//
// 注意：本题与主站 19 题相同： https://leetcode-cn.com/problems/remove-nth-node-from-end-
//of-list/
//
// Related Topics 链表 双指针 👍 93 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := length(head)
	if l < n {
		return head
	}
	var pre = &ListNode{Next: head}
	head = pre
	for i := 0; i < l-n; i++ {
		head = head.Next
	}
	head.Next = head.Next.Next
	return pre.Next
}

func length(head *ListNode) (l int) {
	for ; head != nil; head = head.Next {
		l++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
