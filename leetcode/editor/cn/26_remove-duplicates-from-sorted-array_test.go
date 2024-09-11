package main

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantArray []int
	}{{
		name:      "case1",
		args:      args{[]int{1, 1, 2}},
		want:      2,
		wantArray: []int{1, 2},
	}, {
		name:      "case2",
		args:      args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
		want:      5,
		wantArray: []int{0, 1, 2, 3, 4},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
				return
			}
			if !assertElementEqual(tt.args.nums, tt.wantArray, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums, tt.wantArray)
			}
		})
	}
}
