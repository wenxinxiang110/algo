package main

import (
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		}, {
			name: "test2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		}, {
			name: "test3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		}, {
			name: "test4",
			args: args{
				nums: []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
