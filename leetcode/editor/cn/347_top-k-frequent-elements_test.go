package main

import (
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1：基础场景",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "示例2：单元素数组",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "示例3：重复频率",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "不同频率组合",
			args: args{
				nums: []int{4, 4, 4, 4, 3, 3, 3, 2, 2, 1},
				k:    3,
			},
			want: []int{4, 3, 2},
		},
		{
			name: "负数处理",
			args: args{
				nums: []int{-1, -1, -2, -2, -2},
				k:    2,
			},
			want: []int{-2, -1},
		},
		{
			name: "全元素返回",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "所有元素相同",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "k=1且频率相同",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "所有频率相同",
			args: args{
				nums: []int{1, 1, 2, 2, 3, 3},
				k:    3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "复杂频率组合",
			args: args{
				nums: []int{10, 10, 20, 30, 30, 30},
				k:    2,
			},
			want: []int{30, 10},
		},
		{
			name: "大数测试",
			args: args{
				nums: []int{1000, 1000, 2000, 2000, 2000, 3000},
				k:    2,
			},
			want: []int{2000, 1000},
		},
		{
			name: "零值处理",
			args: args{
				nums: []int{0, 0, 0, 1, 1, 2},
				k:    2,
			},
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)

			// 验证结果长度
			if len(got) != len(tt.want) {
				t.Errorf("topKFrequent() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// 创建期望值的集合
			wantSet := make(map[int]bool)
			for _, v := range tt.want {
				wantSet[v] = true
			}

			// 验证结果中的每个元素都在期望集合中
			for _, v := range got {
				if !wantSet[v] {
					t.Errorf("topKFrequent() = %v, want set %v", got, tt.want)
					return
				}
			}

			// 验证没有重复元素
			gotSet := make(map[int]bool)
			for _, v := range got {
				if gotSet[v] {
					t.Errorf("topKFrequent() contains duplicate element: %v", got)
					return
				}
				gotSet[v] = true
			}
		})
	}
}
