package main

//给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
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
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
//
//
// 注意：本题与主站 206 题相同： https://leetcode-cn.com/problems/reverse-linked-list/
//
// Related Topics 递归 链表 👍 201 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// pre指针记录当前节点的前一个节点地址
	// 初始化的时候，当前节点是头节点，所以pre指针为空
	var pre *ListNode = nil
	for head != nil {
		// pre指针和 head 指针每次往后移一位
		// 反转，所以 next 指到前一个位置
		pre, head, head.Next = head, head.Next, pre
	}
	// 结束循环的时候，当前节点是最后一个节点的下一个，所以 pre 指针指向最后一个节点，也就是反转之后的头节点
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
