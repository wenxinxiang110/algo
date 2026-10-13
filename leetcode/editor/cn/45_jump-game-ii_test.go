package main

import (
	"testing"
)

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "基础情况-最小跳跃次数",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: 2,
		},
		{
			name: "基础情况-一步到达终点",
			args: args{nums: []int{4, 1, 1, 1, 1}},
			want: 1,
		},
		{
			name: "边界情况-单元素数组",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "边界情况-两个元素",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "复杂情况-需要多步跳跃",
			args: args{nums: []int{1, 2, 1, 1, 1}},
			want: 3,
		},
		{
			name: "复杂情况-大跳跃覆盖",
			args: args{nums: []int{2, 1, 3, 1, 1, 1}},
			want: 2,
		},
		{
			name: "边界情况-最小长度数组",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "复杂情况-连续递增",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 3,
		},
		{
			name: "复杂情况-最优路径选择",
			args: args{nums: []int{2, 3, 0, 1, 4}},
			want: 2,
		},
		{
			name: "边界情况-全部为1",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 4,
		},
		{
			name: "复杂情况-需要精确计算",
			args: args{nums: []int{3, 4, 3, 2, 5, 4, 3}},
			want: 3,
		},
		//{
		//	name: "边界情况-第一个元素为0且长度大于1",
		//	args: args{nums: []int{0, 2, 3}},
		//	want: 0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
