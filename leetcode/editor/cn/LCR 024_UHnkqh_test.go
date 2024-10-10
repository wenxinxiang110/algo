package main

import (
	"reflect"
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{{
		name: "test1",
		args: args{head: list.SimpleParse("1->2->3->4->5")},
		want: list.SimpleParse("5->4->3->2->1"),
	}, {
		name: "test2",
		args: args{head: list.SimpleParse("1->2")},
		want: list.SimpleParse("2->1"),
	}, {
		name: "test3",
		args: args{head: nil},
		want: nil,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
