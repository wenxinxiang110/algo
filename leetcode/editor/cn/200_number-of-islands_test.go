package main

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1: 单个岛屿",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
		{
			name: "示例2: 三个岛屿",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		},
		{
			name: "空网格",
			args: args{
				grid: [][]byte{},
			},
			want: 0,
		},
		{
			name: "全是水",
			args: args{
				grid: [][]byte{
					{'0', '0', '0'},
					{'0', '0', '0'},
					{'0', '0', '0'},
				},
			},
			want: 0,
		},
		{
			name: "全是陆地",
			args: args{
				grid: [][]byte{
					{'1', '1', '1'},
					{'1', '1', '1'},
					{'1', '1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "对角线岛屿",
			args: args{
				grid: [][]byte{
					{'1', '0', '0', '0'},
					{'0', '1', '0', '0'},
					{'0', '0', '1', '0'},
					{'0', '0', '0', '1'},
				},
			},
			want: 4,
		},
		{
			name: "单行网格",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '1', '1'},
				},
			},
			want: 2,
		},
		{
			name: "单列网格",
			args: args{
				grid: [][]byte{
					{'1'},
					{'1'},
					{'0'},
					{'1'},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
