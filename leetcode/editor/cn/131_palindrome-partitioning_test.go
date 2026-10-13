package main

import (
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "示例1：aab",
			args: args{s: "aab"},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "示例2：单字符",
			args: args{s: "a"},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "空字符串",
			args: args{s: ""},
			want: [][]string{},
		},
		{
			name: "全相同字符",
			args: args{s: "aaa"},
			want: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "无回文分割",
			args: args{s: "abc"},
			want: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "复杂回文模式",
			args: args{s: "aabb"},
			want: [][]string{
				{"a", "a", "b", "b"},
				{"a", "a", "bb"},
				{"aa", "b", "b"},
				{"aa", "bb"},
			},
		},
		{
			name: "长回文字符串",
			args: args{s: "racecar"},
			want: [][]string{
				{"r", "a", "c", "e", "c", "a", "r"},
				{"r", "a", "cec", "a", "r"},
				{"r", "aceca", "r"},
				{"racecar"},
			},
		},
		{
			name: "包含多个回文",
			args: args{s: "noon"},
			want: [][]string{
				{"n", "o", "o", "n"},
				{"n", "oo", "n"},
				//{"no", "o", "n"},
				//{"n", "oon"},®
				{"noon"},
			},
		},
		{
			name: "混合回文模式",
			args: args{s: "abba"},
			want: [][]string{
				{"a", "b", "b", "a"},
				{"a", "bb", "a"},
				{"abba"},
			},
		},
		{
			name: "边界情况：最大长度字符串",
			args: args{s: "abcdefghijklmnop"},
			want: [][]string{
				{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.s)

			// 验证结果长度
			if len(got) != len(tt.want) {
				t.Errorf("partition() returned %d partitions, want %d", len(got), len(tt.want))
				return
			}

			// 验证每个分割方案
			found := make(map[string]bool)
			for _, partition := range got {
				key := ""
				for _, part := range partition {
					key += part + "|"
				}
				found[key] = true
			}

			for _, expectedPartition := range tt.want {
				key := ""
				for _, part := range expectedPartition {
					key += part + "|"
				}
				if !found[key] {
					t.Errorf("partition() missing expected partition: %v", expectedPartition)
				}
			}

			// 验证没有多余的分割方案
			if len(found) > len(tt.want) {
				t.Errorf("partition() returned extra partitions: got %d, want %d", len(found), len(tt.want))
			}
		})
	}
}

// 辅助函数：验证字符串是否为回文
func isPalindrome4Test(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 验证所有分割方案中的子串都是回文
func Test_partition_all_palindromes(t *testing.T) {
	testCases := []string{"aab", "aaa", "racecar", "noon", "abba"}

	for _, s := range testCases {
		t.Run(s, func(t *testing.T) {
			result := partition(s)

			for _, partition := range result {
				for _, part := range partition {
					if !isPalindrome4Test(part) {
						t.Errorf("partition() returned non-palindrome substring '%s' in partition %v", part, partition)
					}
				}
			}
		})
	}
}
