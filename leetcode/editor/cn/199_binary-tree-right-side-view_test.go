package main

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantView []int
	}{
		{
			"empty",
			args{
				root: nil,
			},
			nil,
		}, {
			"one",
			args{
				root: &TreeNode{
					Val: 1,
				},
			},
			[]int{1},
		}, {
			"two",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			[]int{1, 2},
		}, {
			name: "three",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			wantView: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotView := rightSideView(tt.args.root); !reflect.DeepEqual(gotView, tt.wantView) {
				t.Errorf("rightSideView() = %v, want %v", gotView, tt.wantView)
			}
		})
	}
}
