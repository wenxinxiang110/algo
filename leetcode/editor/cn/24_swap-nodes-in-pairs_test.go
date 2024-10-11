package main

import (
	"reflect"
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: list.SimpleParse("1->2->3->4"),
			},
			want: list.SimpleParse("2->1->4->3"),
		}, {
			name: "test2",
			args: args{
				head: nil,
			},
			want: nil,
		}, {
			name: "test3",
			args: args{
				head: list.SimpleParse("1"),
			},
			want: list.SimpleParse("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
