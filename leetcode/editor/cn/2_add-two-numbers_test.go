package main

import (
	"testing"

	"github.com/NothingXiang/algo/leetcode/list_node"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		}, {
			name: "one",
			args: args{
				l1: list_node.SimpleParse("1->2->3"),
				l2: nil,
			},
			want: list_node.SimpleParse("1->2->3"),
		}, {
			name: "normal",
			args: args{
				l1: list_node.SimpleParse("1->2->3"),
				l2: list_node.SimpleParse("4->5->6"),
			},
			want: list_node.SimpleParse("5->7->9"),
		}, {
			name: "位数不同",
			args: args{
				l1: list_node.SimpleParse("1->2->3"),
				l2: list_node.SimpleParse("4->5"),
			},
			want: list_node.SimpleParse("5->7->3"),
		}, {
			name: "进位",
			args: args{
				l1: list_node.SimpleParse("5->5->3"),
				l2: list_node.SimpleParse("4->5"),
			},
			want: list_node.SimpleParse("9->0->4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !list_node.Equal(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
