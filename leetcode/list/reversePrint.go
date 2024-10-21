package list

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]
func reversePrint(head *ListNode) []int {
	// 计算list长度
	var it = head
	var count = 0
	for it != nil {
		count++
		it = it.Next
	}

	res := make([]int, count)
	cur := count - 1
	for head != nil {
		res[cur] = head.Val
		cur--
		head = head.Next
	}
	return res
}

// 原地反转链表
func reverseList(head *ListNode) *ListNode {
	// 这几个指针用于保存被反转之前的链表顺序
	var (
		pre  *ListNode = nil
		cur            = head
		next *ListNode = nil
	)
	for cur != nil {
		// 先保存位置的原始信息
		next = cur.Next
		// 反转
		cur.Next = pre
		// 保存pre指针
		pre = cur
		// 更新cur到下一个位置
		cur = next
	}
	return pre
}
