package main

import (
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1: 输出4",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "示例2: 输出-1（不可能腐烂）",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			want: -1,
		},
		{
			name: "示例3: 输出0（没有新鲜橘子）",
			args: args{
				grid: [][]int{
					{0, 2},
				},
			},
			want: 0,
		},
		{
			name: "空网格",
			args: args{
				grid: [][]int{},
			},
			want: 0,
		},
		{
			name: "单行网格",
			args: args{
				grid: [][]int{
					{2, 1, 0, 1},
				},
			},
			want: -1,
		},
		{
			name: "单列网格",
			args: args{
				grid: [][]int{
					{2},
					{1},
					{0},
					{1},
				},
			},
			want: -1,
		},
		{
			name: "全是腐烂橘子",
			args: args{
				grid: [][]int{
					{2, 2, 2},
					{2, 2, 2},
				},
			},
			want: 0,
		},
		{
			name: "全是新鲜橘子（没有腐烂源）",
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			want: -1,
		},
		{
			name: "全是空单元格",
			args: args{
				grid: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "多个腐烂源同时开始",
			args: args{
				grid: [][]int{
					{2, 0, 1},
					{0, 1, 0},
					{1, 0, 2},
				},
			},
			want: -1,
		},
		{
			//	[[1,2]]
			name: "示例1",
			args: args{
				grid: [][]int{
					{1, 2},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
