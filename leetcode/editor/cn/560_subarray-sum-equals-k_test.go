package main

import (
	"testing"
)

func Test_subarrayViolent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "empty",
			args: args{
				nums: []int{},
				k:    0,
			},
			wantCount: 0,
		}, {
			name: "single",
			args: args{
				nums: []int{1},
				k:    1,
			},
			wantCount: 1,
		}, {
			name: "multiple",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			wantCount: 2,
		}, {
			name: "multiple",

			args: args{
				nums: []int{3, 4, 7, 2, -3, 1, 4, 2},
				k:    7,
			},
			wantCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := subarrayViolent(tt.args.nums, tt.args.k); gotCount != tt.wantCount {
				t.Errorf("subarrayViolent() = %v, want %v", gotCount, tt.wantCount)
			}
			// test prefix sum
			if gotCount := subarrayPrefixSumV2(tt.args.nums, tt.args.k); gotCount != tt.wantCount {
				t.Errorf("subarrayPrefixSum() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
