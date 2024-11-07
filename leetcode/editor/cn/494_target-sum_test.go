package main

import (
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 1, 1, 1, 1}, target: 3},
			want: 5,
		},
		{
			name: "test2",
			args: args{nums: []int{1}, target: 1},
			want: 1,
		},
		{
			name: "test3",
			args: args{nums: []int{1}, target: 2},
			want: 0,
		},
		{
			name: "test4",
			args: args{nums: []int{1000}, target: -1000},
			want: 1,
		},
		{
			name: "test5",
			args: args{nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 1}, target: 1},
			want: 256,
		},
		{
			name: "test6",
			args: args{nums: []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}, target: 0},
			want: 0,
		},
		{
			name: "test7",
			args: args{nums: []int{9, 7, 0, 3, 9, 8, 6, 5, 7, 6}, target: 2},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
