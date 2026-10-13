package main

import (
	"reflect"
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_reverseListPreK(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name        string
		args        args
		wantNewHead *ListNode
	}{
		{
			name: "empty",
			args: args{
				head: nil,
				k:    2,
			},
			wantNewHead: nil,
		}, {
			name: "test1",
			args: args{
				head: list.SimpleParse("1->2->3->4->5"),
				k:    2,
			},
			wantNewHead: list.SimpleParse("2->1"),
		}, {
			name: "test2",
			args: args{
				head: list.SimpleParse("1->2->3->4->5"),
				k:    6,
			},
			wantNewHead: list.SimpleParse("5->4->3->2->1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewHead, _ := reverseListPreK(tt.args.head, tt.args.k)
			if !reflect.DeepEqual(gotNewHead, tt.wantNewHead) {
				t.Errorf("reverseListPreK() gotNewHead = %v, want %v", gotNewHead, tt.wantNewHead)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{
				head: nil,
				k:    2,
			},
			want: nil,
		}, {
			name: "test1",
			args: args{
				head: list.SimpleParse("1->2->3->4->5"),
				k:    2,
			},
			want: list.SimpleParse("2->1->4->3->5"),
		}, {
			name: "test2",
			args: args{
				head: list.SimpleParse("1->2->3->4->5"),
				k:    3,
			},
			want: list.SimpleParse("3->2->1->4->5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
