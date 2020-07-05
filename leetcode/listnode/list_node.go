package listnode

import (
	"log"
	"strconv"
	"strings"
)

/**
 * todo: 今日总结
 * 1. 添加一个假的头节点，这样当遇到真正的头节点时就不用特殊处理了；
 * 2. 一边遍历一边修改的时候，至少需要两个指针处理，毕竟链表不能访问到上一个节点
 *
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

/**
 * 对@ParseListNode方法的简单封装
 *
 * param: string listExp
 * return: *ListNode
 */
func SimpleParse(listExp string) *ListNode {
	return ParseListNode(listExp, "->")
}

/**
 * 根据表达式和分割符生成链表
 * 例如：
 * 1->1->2->3->3
 * 1->1->2
 *
 * param: string exp 链表表达式
 * param: string spt 表达式的分割符
 * return: *ListNode 生成的链表的头节点
 */
func ParseListNode(exp, spt string) *ListNode {

	// 一个没有用的假的头节点，这样遇到真正的第一个节点的时候就不用特殊处理了
	head := NewListNode(0)

	// 当前处理的节点
	var cur = head

	// 节点值
	vals := strings.Split(exp, spt)

	for _, val := range vals {
		value, err := strconv.Atoi(val)

		// todo: 出错处理暂时不弄，直接返回nil
		if err != nil {
			log.Fatal(err)
			return nil
		}

		cur.Next = NewListNode(value)
		cur = cur.Next
	}

	return head.Next
}

/**
 * 判断两个链表是否相等
 *
 * param: *ListNode first 第一个链表的头节点
 * param: *ListNode second 第二个链表的头节点
 * return: bool
 */
func Equal(first, second *ListNode) bool {
	for first != nil && second != nil {

		if first.Val != second.Val {
			return false
		}

		first = first.Next
		second = second.Next

	}

	if first != nil || second != nil {
		return false
	}

	return true
}
