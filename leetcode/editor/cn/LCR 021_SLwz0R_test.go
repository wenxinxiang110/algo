package main

import (
	"reflect"
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: list.SimpleParse("1->2->3->4->5"),
				n:    2,
			},
			want: list.SimpleParse("1->2->3->5"),
		},
		{
			name: "case2",
			args: args{
				head: list.SimpleParse("1"),
				n:    1,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				head: list.SimpleParse("1->2"),
				n:    1,
			},
			want: list.SimpleParse("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
