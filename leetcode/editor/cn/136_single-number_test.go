package main

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantSingle int
	}{
		{
			name:       "示例1: [2,2,1]",
			args:       args{nums: []int{2, 2, 1}},
			wantSingle: 1,
		},
		{
			name:       "示例2: [4,1,2,1,2]",
			args:       args{nums: []int{4, 1, 2, 1, 2}},
			wantSingle: 4,
		},
		{
			name:       "示例3: [1]",
			args:       args{nums: []int{1}},
			wantSingle: 1,
		},
		{
			name:       "边界情况: 最小长度数组",
			args:       args{nums: []int{5}},
			wantSingle: 5,
		},
		{
			name:       "负数测试: [-1,-1,-2]",
			args:       args{nums: []int{-1, -1, -2}},
			wantSingle: -2,
		},
		{
			name:       "混合正负数: [3,-3,3,5,-3]",
			args:       args{nums: []int{3, -3, 3, 5, -3}},
			wantSingle: 5,
		},
		{
			name:       "零值测试: [0,1,0]",
			args:       args{nums: []int{0, 1, 0}},
			wantSingle: 1,
		},
		{
			name:       "全零数组: [0,0,0,0,1]",
			args:       args{nums: []int{0, 0, 0, 0, 1}},
			wantSingle: 1,
		},
		{
			name:       "大数测试: [10000,20000,10000]",
			args:       args{nums: []int{10000, 20000, 10000}},
			wantSingle: 20000,
		},
		{
			name:       "复杂场景: 多个重复对",
			args:       args{nums: []int{1, 2, 3, 4, 5, 1, 2, 3, 4}},
			wantSingle: 5,
		},
		{
			name:       "边界值: 最大值测试",
			args:       args{nums: []int{30000, 20000, 30000}},
			wantSingle: 20000,
		},
		{
			name:       "边界值: 最小值测试",
			args:       args{nums: []int{-30000, -20000, -30000}},
			wantSingle: -20000,
		},
		{
			name:       "交替模式: [a,b,a,c,b,d,c]",
			args:       args{nums: []int{10, 20, 10, 30, 20, 40, 30}},
			wantSingle: 40,
		},
		{
			name:       "对称模式: [1,2,3,2,1]",
			args:       args{nums: []int{1, 2, 3, 2, 1}},
			wantSingle: 3,
		},
		{
			name:       "长数组测试",
			args:       args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			wantSingle: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSingle := singleNumber(tt.args.nums); gotSingle != tt.wantSingle {
				t.Errorf("singleNumber() = %v, want %v", gotSingle, tt.wantSingle)
			}
		})
	}
}
