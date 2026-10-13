package main

import (
	"strconv"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "empty arrays",
			args: args{
				preorder: []int{},
				inorder:  []int{},
			},
			want: nil,
		},
		{
			name: "single node",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: &TreeNode{
				Val: -1,
			},
		},
		{
			name: "example 1 - balanced tree",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "left skewed tree",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{3, 2, 1},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "right skewed tree",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{1, 2, 3},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "complex tree",
			args: args{
				preorder: []int{1, 2, 4, 5, 3, 6, 7},
				inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.preorder, tt.args.inorder)
			if !isTreeEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", treeToString(got), treeToString(tt.want))
			}
		})
	}
}

// isTreeEqual 检查两棵树是否相等
func isTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isTreeEqual(t1.Left, t2.Left) && isTreeEqual(t1.Right, t2.Right)
}

// treeToString 将树转换为字符串表示，用于错误信息
func treeToString(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	var result string
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result += "null,"
			continue
		}

		result += strconv.Itoa(node.Val) + ","
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// 移除末尾多余的逗号
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
