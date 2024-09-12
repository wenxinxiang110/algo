package main

import (
	"reflect"
	"testing"
)

func Test_spiralArray(t *testing.T) {
	type args struct {
		array [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{array: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name: "test2",
			args: args{array: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}, {
			name: "test3",
			args: args{array: [][]int{{2, 3}}},
			want: []int{2, 3},
		}, {
			name: "test4",
			args: args{array: [][]int{{6, 9, 7}}},
			want: []int{6, 9, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
