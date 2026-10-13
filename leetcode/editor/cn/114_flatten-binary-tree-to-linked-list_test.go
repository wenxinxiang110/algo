package main

import (
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "empty tree",
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			name: "single node",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "simple left subtree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "simple right subtree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "complex tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制原始树结构，避免测试间的相互影响
			originalTree := copyTree(tt.args.root)

			flatten(originalTree)

			// 验证展开后的链表结构
			if !isLinkedListEqual(originalTree, tt.want) {
				t.Errorf("flatten() result incorrect")
			}
		})
	}
}

// copyTree 复制二叉树
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

// isLinkedListEqual 检查两个链表是否相等
func isLinkedListEqual(l1, l2 *TreeNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		// 检查左子树是否为nil
		if l1.Left != nil {
			return false
		}
		l1 = l1.Right
		l2 = l2.Right
	}
	return l1 == nil && l2 == nil
}
