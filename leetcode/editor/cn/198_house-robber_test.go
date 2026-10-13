package main

import (
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "边界情况-空数组",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "边界情况-单元素",
			args: args{nums: []int{5}},
			want: 5,
		},
		{
			name: "基础情况-两个元素",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "基础情况-三个元素",
			args: args{nums: []int{1, 2, 3}},
			want: 4,
		},
		{
			name: "经典示例-四个元素",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "经典示例-五个元素",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "复杂情况-多个元素",
			args: args{nums: []int{6, 7, 1, 3, 8, 2, 4}},
			want: 19,
		},
		{
			name: "特殊情况-所有房子价值相同",
			args: args{nums: []int{5, 5, 5, 5, 5}},
			want: 15,
		},
		{
			name: "特殊情况-递增序列",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 9,
		},
		{
			name: "特殊情况-递减序列",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 9,
		},
		{
			name: "复杂情况-大数值",
			args: args{nums: []int{100, 50, 200, 150, 300}},
			want: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
