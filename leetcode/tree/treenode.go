package tree

import (
	"container/list"
)

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Visitor node will not equal nil
type Visitor = func(node *TreeNode)

// PreorderTraversal 前序遍历
func (root *TreeNode) PreorderTraversal(visit Visitor) {
	// 1. 确定终止条件
	if root == nil {
		return
	}
	//确定单层递归的逻辑
	visit(root)
	root.Left.PreorderTraversal(visit)
	root.Right.PreorderTraversal(visit)
}

// PreorderTraversal 前序遍历非递归实现
func (root *TreeNode) preorderTraversalIter(visit Visitor) {
	var stack = list.New()
	stack.PushBack(root)
	for stack.Back() != nil {
		top, _ := stack.Remove(stack.Back()).(*TreeNode)
		visit(top)
		if top.Left != nil {
			stack.PushBack(top.Left)
		} else {
			stack.PushBack(top.Right)
		}
	}

}

// InorderTraversal 中序遍历
func (root *TreeNode) InorderTraversal(visit Visitor) {
	// 1. 确定终止条件
	if root == nil {
		return
	}
	//确定单层递归的逻辑
	root.Left.InorderTraversal(visit)
	visit(root)
	root.Right.InorderTraversal(visit)
}

// PostorderTraversal 后序遍历
func (root *TreeNode) PostorderTraversal(visit Visitor) {
	// 1. 确定终止条件
	if root == nil {
		return
	}
	//确定单层递归的逻辑
	root.Left.PostorderTraversal(visit)
	root.Right.PostorderTraversal(visit)
	visit(root)
}
