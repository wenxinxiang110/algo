package main

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name        string
		args        args
		wantProfile int
	}{
		{
			name:        "基础情况-有利润",
			args:        args{prices: []int{7, 1, 5, 3, 6, 4}},
			wantProfile: 5,
		},
		{
			name:        "边界情况-没有利润",
			args:        args{prices: []int{7, 6, 4, 3, 1}},
			wantProfile: 0,
		},
		{
			name:        "边界情况-空数组",
			args:        args{prices: []int{}},
			wantProfile: 0,
		},
		{
			name:        "边界情况-单元素数组",
			args:        args{prices: []int{5}},
			wantProfile: 0,
		},
		{
			name:        "复杂情况-价格波动较大",
			args:        args{prices: []int{2, 10, 1, 8, 15, 3}},
			wantProfile: 14,
		},
		{
			name:        "边界情况-价格全部相同",
			args:        args{prices: []int{5, 5, 5, 5, 5}},
			wantProfile: 0,
		},
		{
			name:        "复杂情况-多段利润",
			args:        args{prices: []int{3, 2, 6, 5, 0, 3}},
			wantProfile: 4,
		},
		{
			name:        "边界情况-最小长度数组",
			args:        args{prices: []int{1, 2}},
			wantProfile: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProfile := maxProfit(tt.args.prices); gotProfile != tt.wantProfile {
				t.Errorf("maxProfit() = %v, want %v", gotProfile, tt.wantProfile)
			}
		})
	}
}
