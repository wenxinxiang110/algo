package main

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			wantResult: [][]int{
				{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
			},
		}, {
			name: "test2",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			wantResult: [][]int{
				{2, 2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("fourSum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
