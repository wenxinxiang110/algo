package main

import (
	"testing"

	"github.com/NothingXiang/algo/leetcode/list"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				head: nil,
			},
			want: true,
		}, {
			name: "test1",
			args: args{
				head: list.SimpleParse("1->2->2->1"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
