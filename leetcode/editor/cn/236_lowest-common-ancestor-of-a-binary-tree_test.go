package main

import (
	"testing"
)

// 辅助函数：根据值查找节点
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

// 构建示例1的二叉树：root = [3,5,1,6,2,0,8,null,null,7,4]
func buildExample1Tree() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	return root
}

// 构建示例2的二叉树：与示例1相同结构
func buildExample2Tree() *TreeNode {
	return buildExample1Tree()
}

// 构建示例3的二叉树：root = [1,2]
func buildExample3Tree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	return root
}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int // 期望的公共祖先节点值
	}{
		{
			name: "示例1: p=5, q=1",
			args: args{
				root: buildExample1Tree(),
				p:    findNode(buildExample1Tree(), 5),
				q:    findNode(buildExample1Tree(), 1),
			},
			want: 3,
		},
		{
			name: "示例2: p=5, q=4",
			args: args{
				root: buildExample2Tree(),
				p:    findNode(buildExample2Tree(), 5),
				q:    findNode(buildExample2Tree(), 4),
			},
			want: 5,
		},
		{
			name: "示例3: p=1, q=2",
			args: args{
				root: buildExample3Tree(),
				p:    findNode(buildExample3Tree(), 1),
				q:    findNode(buildExample3Tree(), 2),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q)
			if got == nil {
				t.Errorf("lowestCommonAncestor() returned nil, want %d", tt.want)
				return
			}
			if got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %d, want %d", got.Val, tt.want)
			}
		})
	}
}
