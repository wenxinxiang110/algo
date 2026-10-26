package main

import (
	"reflect"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 基本中间匹配
		{name: "basic mid", args: args{haystack: "hello", needle: "ll"}, want: 2},
		// 从下标 0 开始匹配
		{name: "start from 0", args: args{haystack: "sadbutsad", needle: "sad"}, want: 0},
		// 在末尾匹配
		{name: "at end", args: args{haystack: "hello", needle: "lo"}, want: 3},
		// 单字符匹配
		{name: "single char match", args: args{haystack: "a", needle: "a"}, want: 0},
		// 单字符无匹配
		{name: "single char no match", args: args{haystack: "aaa", needle: "b"}, want: -1},
		// needle 比 haystack 长
		{name: "needle longer than haystack", args: args{haystack: "abc", needle: "abcd"}, want: -1},
		// 两串完全相等
		{name: "equal strings", args: args{haystack: "abc", needle: "abc"}, want: 0},
		// 空 needle：按题意返回 0
		{name: "empty needle", args: args{haystack: "abc", needle: ""}, want: 0},
		// 空 haystack 且 needle 非空
		{name: "empty haystack non-empty needle", args: args{haystack: "", needle: "a"}, want: -1},
		// 完全无匹配
		{name: "no match", args: args{haystack: "leetcode", needle: "leeto"}, want: -1},
		// 经典 KMP 重叠回退场景
		{name: "kmp overlap backtrack", args: args{haystack: "abxabcabcaby", needle: "abcaby"}, want: 6},
		// mississippi 中匹配 issip
		{name: "mississippi issip", args: args{haystack: "mississippi", needle: "issip"}, want: 4},
		// 多次出现时返回第一个
		{name: "first of multiple", args: args{haystack: "mississippi", needle: "issi"}, want: 1},
		// 重复模式 + 前缀重叠
		{name: "repeat pattern", args: args{haystack: "aaabaaab", needle: "aaab"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPrefixTable(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name      string
		args      args
		wantTable []int
	}{
		// TODO: Add test cases.
		// 空串：函数直接返回 nil
		{
			name:      "empty",
			args:      args{s: []byte("")},
			wantTable: nil,
		},
		// 单字符：最长公共前后缀长度固定为 0
		{
			name:      "single char",
			args:      args{s: []byte("a")},
			wantTable: []int{0},
		},
		// 完全无重复字符：每个位置的最长公共前后缀长度均为 0
		{
			name:      "no repeat",
			args:      args{s: []byte("abcdef")},
			wantTable: []int{0, 0, 0, 0, 0, 0},
		},
		// 全部相同：长度递增 0,1,2,...,n-1
		{
			name:      "all same",
			args:      args{s: []byte("aaaa")},
			wantTable: []int{0, 1, 2, 3},
		},
		// 仅前缀局部重复，存在不匹配不回退的情况
		{
			name:      "partial prefix repeat",
			args:      args{s: []byte("aaba")},
			wantTable: []int{0, 1, 0, 1},
		},
		// 嵌套前后缀：需要连续匹配累加
		{
			name:      "nested prefix",
			args:      args{s: []byte("ababc")},
			wantTable: []int{0, 0, 1, 2, 0},
		},
		// 经典 KMP 教学串：包含多次回退
		{
			name:      "abababca",
			args:      args{s: []byte("abababca")},
			wantTable: []int{0, 0, 1, 2, 3, 4, 0, 1},
		},
		// 周期性重复串：前缀表呈现周期性递增
		{
			name:      "agctagcagctagct",
			args:      args{s: []byte("agctagcagctagct")},
			wantTable: []int{0, 0, 0, 0, 1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTable := myPrefixTable(tt.args.s); !reflect.DeepEqual(gotTable, tt.wantTable) {
				t.Errorf("myPrefixTable() = %v, want %v", gotTable, tt.wantTable)
			}
		})
	}
}
