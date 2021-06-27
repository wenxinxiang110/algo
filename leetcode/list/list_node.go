package list

import (
	"bytes"
	"strconv"
	"strings"
)

/**
 * todo: 今日总结
 * 1. 添加一个假的头节点，这样当遇到真正的头节点时就不用特殊处理了；
 * 2. 一边遍历一边修改的时候，至少需要两个指针处理，毕竟链表不能访问到上一个节点
 *
 */

const (
	// 链表表达式默认分割符
	defaultSep = "->"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}

	res := new(bytes.Buffer)
	res.WriteString(strconv.Itoa(l.Val))
	cur := l.Next
	for cur != nil {
		res.WriteString(defaultSep)
		res.WriteString(strconv.Itoa(cur.Val))
		cur = cur.Next
	}

	return res.String()
}

// NewListNode 构造函数
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// SimpleParse 对ParseListNode方法的简单封装，使用默认分隔符解析链表
// param: listExp 链表的表达式
// return: 头结点
func SimpleParse(listExp string) *ListNode {
	return ParseListNode(listExp, defaultSep)
}

// ParseListNode 根据表达式和分割符生成链表
// 例如：
// 1->1->2->3->3
// 1->1->2
// param: exp 链表表达式
// param: spt 表达式的分割符
// return: 生成的链表的头节点
func ParseListNode(exp, spt string) *ListNode {
	var (
		// 一个没有用的假的头节点，这样遇到真正的第一个节点的时候就不用特殊处理了
		head = NewListNode(0)
		// 当前处理的节点
		cur = head
		// 节点值
		val = strings.Split(exp, spt)
	)
	for _, v := range val {
		if v == "" {
			continue
		}
		value, _ := strconv.Atoi(v)
		cur.Next = NewListNode(value)
		cur = cur.Next
	}

	return head.Next
}

// Equal 判断两个链表是否相等
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
