package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1：k=3,n=7 唯一组合",
			args: args{k: 3, n: 7},
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "示例2：k=3,n=9 多组合",
			args: args{k: 3, n: 9},
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "示例3：k=4,n=1 小于最小和10",
			args: args{k: 4, n: 1},
			want: [][]int{},
		},
		{
			name: "k=2,n=3 最小和",
			args: args{k: 2, n: 3},
			want: [][]int{{1, 2}},
		},
		{
			name: "k=2,n=5 多组合",
			args: args{k: 2, n: 5},
			want: [][]int{{1, 4}, {2, 3}},
		},
		{
			name: "k=2,n=17 最大和",
			args: args{k: 2, n: 17},
			want: [][]int{{8, 9}},
		},
		{
			name: "k=2,n=18 超过最大和",
			args: args{k: 2, n: 18},
			want: [][]int{},
		},
		{
			name: "k=3,n=2 小于最小和6",
			args: args{k: 3, n: 2},
			want: [][]int{},
		},
		{
			name: "k=3,n=6 最小和",
			args: args{k: 3, n: 6},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "k=3,n=24 最大和",
			args: args{k: 3, n: 24},
			want: [][]int{{7, 8, 9}},
		},
		{
			name: "k=3,n=25 超过最大和",
			args: args{k: 3, n: 25},
			want: [][]int{},
		},
		{
			name: "k=3,n=15 八种组合",
			args: args{k: 3, n: 15},
			want: [][]int{
				{1, 5, 9}, {1, 6, 8}, {2, 4, 9}, {2, 5, 8},
				{2, 6, 7}, {3, 4, 8}, {3, 5, 7}, {4, 5, 6},
			},
		},
		{
			name: "k=4,n=10 最小和",
			args: args{k: 4, n: 10},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "k=5,n=15 唯一组合",
			args: args{k: 5, n: 15},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "k=8,n=36 取前八个数字",
			args: args{k: 8, n: 36},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
		},
		{
			name: "k=9,n=45 取全部1到9",
			args: args{k: 9, n: 45},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name: "k=9,n=44 无解",
			args: args{k: 9, n: 44},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.args.k, tt.args.n)
			if !reflect.DeepEqual(normalizeCombinationSum3(got), normalizeCombinationSum3(tt.want)) {
				t.Errorf("combinationSum3(%d, %d) = %v, want %v", tt.args.k, tt.args.n, got, tt.want)
			}
		})
	}
}

// TestCombinationSum3Invariants 校验结果满足题面约束：k个数、和为n、
// 仅使用1到9且严格递增（每个数字最多一次）、组合不重复。
func TestCombinationSum3Invariants(t *testing.T) {
	tests := []struct {
		name      string
		k         int
		n         int
		wantCount int
	}{
		{name: "k=2,n=6", k: 2, n: 6, wantCount: 2},
		{name: "k=3,n=12", k: 3, n: 12, wantCount: 7},
		{name: "k=4,n=20", k: 4, n: 20, wantCount: 12},
		{name: "k=5,n=25", k: 5, n: 25, wantCount: 12},
		{name: "k=6,n=30", k: 6, n: 30, wantCount: 8},
		{name: "k=7,n=35", k: 7, n: 35, wantCount: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			if len(got) != tt.wantCount {
				t.Fatalf("combinationSum3(%d, %d) 返回 %d 个组合，期望 %d 个", tt.k, tt.n, len(got), tt.wantCount)
			}
			seen := make(map[string]bool)
			for _, comb := range got {
				if len(comb) != tt.k {
					t.Fatalf("combinationSum3(%d, %d) 返回组合 %v 长度为 %d，期望 %d", tt.k, tt.n, comb, len(comb), tt.k)
				}
				sum := 0
				for i, num := range comb {
					if num < 1 || num > 9 {
						t.Fatalf("combinationSum3(%d, %d) 返回组合 %v 含非法数字 %d，只允许1到9", tt.k, tt.n, comb, num)
					}
					if i > 0 && comb[i-1] >= num {
						t.Fatalf("combinationSum3(%d, %d) 返回组合 %v 未严格递增，数字被重复使用", tt.k, tt.n, comb)
					}
					sum += num
				}
				if sum != tt.n {
					t.Fatalf("combinationSum3(%d, %d) 返回组合 %v 和为 %d，期望 %d", tt.k, tt.n, comb, sum, tt.n)
				}
				key := fmt.Sprint(comb)
				if seen[key] {
					t.Fatalf("combinationSum3(%d, %d) 返回重复组合 %v", tt.k, tt.n, comb)
				}
				seen[key] = true
			}
		})
	}
}

// normalizeCombinationSum3 对结果做规范化排序，忽略组合之间的顺序，
// 便于按集合语义比较（题面允许以任意顺序返回组合）。
func normalizeCombinationSum3(comb [][]int) [][]int {
	normalized := make([][]int, 0, len(comb))
	for _, c := range comb {
		item := make([]int, len(c))
		copy(item, c)
		sort.Ints(item)
		normalized = append(normalized, item)
	}
	sort.Slice(normalized, func(i, j int) bool {
		if len(normalized[i]) != len(normalized[j]) {
			return len(normalized[i]) < len(normalized[j])
		}
		for x := 0; x < len(normalized[i]); x++ {
			if normalized[i][x] != normalized[j][x] {
				return normalized[i][x] < normalized[j][x]
			}
		}
		return false
	})
	return normalized
}
