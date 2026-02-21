package main

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			// 中序遍历打印结果，比较是否相同
			gotResult := inorderTraversal(got)
			if !reflect.DeepEqual(gotResult, tt.args.nums) {
				t.Errorf("sortedArrayToBST() = %v, want %v", gotResult, tt.args.nums)
			}
		})
	}
}
