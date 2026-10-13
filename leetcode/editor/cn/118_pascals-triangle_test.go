package main

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "边界情况-0行",
			args: args{numRows: 0},
			want: [][]int{},
		},
		{
			name: "边界情况-1行",
			args: args{numRows: 1},
			want: [][]int{{1}},
		},
		{
			name: "基础情况-2行",
			args: args{numRows: 2},
			want: [][]int{{1}, {1, 1}},
		},
		{
			name: "基础情况-3行",
			args: args{numRows: 3},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name: "基础情况-4行",
			args: args{numRows: 4},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			name: "基础情况-5行",
			args: args{numRows: 5},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "复杂情况-6行",
			args: args{numRows: 6},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
		},
		{
			name: "复杂情况-7行",
			args: args{numRows: 7},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}},
		},
		//{
		//	name: "边界情况-负数行数",
		//	args: args{numRows: -1},
		//	want: nil,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
