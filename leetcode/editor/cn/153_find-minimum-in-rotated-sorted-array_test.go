package main

import (
	"testing"
)

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1：旋转数组，最小元素在中间",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "示例2：旋转数组，最小元素在中间偏后",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "示例3：未旋转的升序数组",
			args: args{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
		{
			name: "单元素数组",
			args: args{
				nums: []int{5},
			},
			want: 5,
		},
		{
			name: "双元素数组，已旋转",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "双元素数组，未旋转",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "最小元素在开头",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 6, 7},
			},
			want: 0,
		},
		{
			name: "最小元素在末尾",
			args: args{
				nums: []int{2, 4, 5, 6, 7, 0, 1},
			},
			want: 0,
		},
		{
			name: "复杂旋转场景",
			args: args{
				nums: []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 0,
		},
		{
			name: "负数和正数混合",
			args: args{
				nums: []int{-3, -2, -1, 0, 1, 2, -5, -4},
			},
			want: -5,
		},
		{
			name: "全负数数组",
			args: args{
				nums: []int{-5, -4, -3, -2, -1},
			},
			want: -5,
		},
		{
			name: "边界值测试：最小整数",
			args: args{
				nums: []int{-5000, -4999, -4998},
			},
			want: -5000,
		},
		{
			name: "边界值测试：最大整数",
			args: args{
				nums: []int{4998, 4999, 5000},
			},
			want: 4998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
