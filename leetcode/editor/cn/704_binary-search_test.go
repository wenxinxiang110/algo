package main

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{5},
				target: -5,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 5},
				target: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "目标在中间存在",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "目标不存在返回-1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "单元素命中",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 0,
		},
		{
			name: "单元素未命中",
			args: args{
				nums:   []int{5},
				target: -5,
			},
			want: -1,
		},
		{
			name: "目标在末尾",
			args: args{
				nums:   []int{2, 5},
				target: 5,
			},
			want: 1,
		},
		{
			name: "目标在开头",
			args: args{
				nums:   []int{2, 5},
				target: 2,
			},
			want: 0,
		},
		{
			name: "较多元素的边界场景",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				target: 1,
			},
			want: 0,
		},
		{
			name: "较多元素的边界场景2",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				target: 10,
			},
			want: 9,
		},
		{
			name: "目标不存在于较大数组",
			args: args{
				nums:   []int{1, 3, 5, 7, 9, 11, 13, 15},
				target: 8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
