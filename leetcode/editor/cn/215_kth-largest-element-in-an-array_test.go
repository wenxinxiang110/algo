package main

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			nums: []int{2, 1},
			k:    1,
			want: 2,
		},
		{
			nums: []int{2, 1},
			k:    2,
			want: 1,
		},
		{
			nums: []int{5, 5, 5, 5, 5},
			k:    3,
			want: 5,
		},
		{
			nums: []int{-1, -2, -3, -4, -5},
			k:    2,
			want: -2,
		},
	}

	for _, tt := range tests {
		got := findKthLargest(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}

func TestFindKthLargest_EdgeCases(t *testing.T) {
	// 测试边界情况
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{10, 20, 30, 40, 50},
			k:    1,
			want: 50, // 最大的元素
		},
		{
			nums: []int{10, 20, 30, 40, 50},
			k:    5,
			want: 10, // 最小的元素
		},
		{
			nums: []int{1, 3, 2, 5, 4},
			k:    3,
			want: 3, // 中间值
		},
	}

	for _, tc := range testCases {
		got := findKthLargest(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("findKthLargest(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}
