package main

import (
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{},
			want: false,
		}, {
			name: "single",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 1,
			},
			want: true,
		}, {
			name: "two",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
				target: 1,
			},
			want: true,
		}, {
			name: "target not found",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
				target: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
