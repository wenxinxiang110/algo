package main

import (
	"testing"
)

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "基础情况-可以到达终点",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "基础情况-无法到达终点",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "边界情况-单元素数组",
			args: args{nums: []int{0}},
			want: true,
		},
		{
			name: "边界情况-两个元素且可到达",
			args: args{nums: []int{1, 0}},
			want: true,
		},
		{
			name: "边界情况-两个元素且不可到达",
			args: args{nums: []int{0, 1}},
			want: false,
		},
		{
			name: "复杂情况-中间有0但可以跳过",
			args: args{nums: []int{2, 0, 1, 1, 4}},
			want: true,
		},
		{
			name: "复杂情况-中间有0且无法跳过",
			args: args{nums: []int{1, 0, 2, 0, 4}},
			want: false,
		},
		{
			name: "边界情况-全部为0",
			args: args{nums: []int{0, 0, 0, 0}},
			want: false,
		},
		{
			name: "复杂情况-大跳跃可以覆盖",
			args: args{nums: []int{6, 0, 0, 0, 0, 0, 1}},
			want: true,
		},
		{
			name: "边界情况-最小长度数组",
			args: args{nums: []int{1}},
			want: true,
		},
		{
			name: "复杂情况-连续跳跃",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "边界情况-第一个元素为0且长度大于1",
			args: args{nums: []int{0, 2, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
