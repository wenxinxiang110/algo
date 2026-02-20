package main

import (
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
	"github.com/NothingXiang/algo/leetcode/list_node"
)

func Test_sortListTopdown(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"empty",
			args{
				head: nil,
			},
			nil,
		}, {
			"single",
			args{
				head: list.SimpleParse("1"),
			},
			list.SimpleParse("1"),
		}, {
			"奇数",
			args{
				head: list.SimpleParse("2->1->3"),
			},
			list.SimpleParse("1->2->3"),
		}, {
			"偶数",
			args{
				head: list.SimpleParse("4->3->2->1"),
			},
			list.SimpleParse("1->2->3->4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortListTopdown(tt.args.head); !list_node.Equal(got, tt.want) {
				t.Errorf("sortListTopdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
