package main

import (
	"fmt"
)

//给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 回溯 👍 2170 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) (res [][]string) {
	if len(s) == 0 {
		return
	}

	//  table[i][j]表示 s[i:j+1]是否为回文串 0未搜索， -1 不是， 1是回文
	table := make([][]int8, len(s))
	for i := range table {
		table[i] = make([]int8, len(s))
	}

	const (
		notPalindrome = -1
		unsearch      = 0
		palindrome    = 1
	)
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		// 相等的时候肯定是回文
		if i >= j {
			return palindrome
		}
		// 已经搜索过了，直接返回
		if table[i][j] != unsearch {
			return table[i][j]
		}
		// 不是回文
		if s[i] != s[j] {
			table[i][j] = notPalindrome
		} else {
			table[i][j] = isPalindrome(i+1, j-1)
		}
		return table[i][j]
	}

	var split []string
	var backtrace func(depth int)
	backtrace = func(depth int) {
		// 已经找到了处理方案
		if depth == len(s) {
			fmt.Println("找到split:", split, "depth:", depth)
			res = append(res, append([]string(nil), split...))
			return
		}

		for j := depth; j < len(s); j++ {
			if isPalindrome(depth, j) == palindrome {
				split = append(split, s[depth:j+1])
				fmt.Println("回朔之前,split:", split, "depth:", depth, "j:", j)
				backtrace(j + 1)
				split = split[:len(split)-1]
				fmt.Println("回朔之后,split:", split, "depth:", depth, "j:", j)
			}
		}
	}

	backtrace(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
