package main

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				matrix: [][]int{},
			},
			want: nil,
		}, {
			name: "one",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			want: []int{1},
		},
		{
			name: "奇数正方形",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		}, {
			name: "偶数正方形",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: []int{1, 2, 4, 3},
		}, {
			name: "长一点的矩阵",
			args: args{
				matrix: [][]int{
					{1, 2},
				},
			},
			want: []int{1, 2},
		}, {
			//	测试用例:[[6,9,7]]
			name: "长一点的矩阵2",
			args: args{
				matrix: [][]int{
					{6, 9, 7},
				},
			},
			want: []int{6, 9, 7},
		}, {
			name: "编一点的矩阵",
			args: args{
				matrix: [][]int{
					{3},
					{2},
				},
			},
			want: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
