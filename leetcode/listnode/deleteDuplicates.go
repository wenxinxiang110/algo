package listnode

//83. 删除排序链表中的重复元素
//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3
func deleteDuplicates(head *ListNode) *ListNode {

	// 迭代器
	iter := head

	// 当前节点的值
	curVal := 0

	// 上一个节点的地址
	var perNode *ListNode = nil

	if head != nil {
		curVal = head.Val
		perNode = iter
		iter = iter.Next
	}

	for ; iter != nil; iter = iter.Next {

		// 如果是重复节点，跳过重复节点，连接到重复节点的下一个节点
		if iter.Val == curVal {
			perNode.Next = iter.Next
			continue
		}

		// 如果不是重复节点,更新当前值用于判断下一个节点是否重复
		curVal = iter.Val
		perNode = iter

	}

	return head

}
