package main

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
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
			name: "test1",
			args: args{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: []int{3, 3, 5, 5, 6, 7},
		}, {
			name: "test2",
			args: args{nums: []int{1}, k: 1},
			want: []int{1},
		}, {
			name: "test3",
			args: args{nums: []int{1, -1}, k: 1},
			want: []int{1, -1},
		}, {
			name: "test4",
			args: args{nums: []int{7, 2, 4}, k: 2},
			want: []int{7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
