package main

import (
	"reflect"
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 1",
			args: args{
				head: list.SimpleParse("1->2->6->3->4->5->6"),
				val:  6,
			},
			want: list.SimpleParse("1->2->3->4->5"),
		},
		{
			name: "test 2",
			args: args{
				head: list.SimpleParse("1->2->6->3->4->5->6"),
				val:  1,
			},
			want: list.SimpleParse("2->6->3->4->5->6"),
		},
		{
			name: "test 3",
			args: args{
				head: list.SimpleParse("7->7->7->7"),
				val:  7,
			},
			want: nil,
		},
		{
			name: "empty",
			args: args{
				val: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
