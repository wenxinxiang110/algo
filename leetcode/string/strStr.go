package string

import (
	"strings"
)

//28. 实现 strStr()
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
func strStr(haystack string, needle string) int {

	// 特殊处理
	if needle == "" {
		return 0
	}

	if haystack == "" {
		return -1
	}

	// 等价于下面那一段循环，时间复杂度O(n*m)
	return strings.Index(haystack, needle)
	/*
		for i := 0; i < len(haystack)-len(needle); i++ {
			if haystack[i] != needle[0] {
				continue
			}

			for j := 0; j < len(needle); j++ {
				if needle[j] != haystack[i+j] {
					break
				}

				if j == len(needle)-1 {
					return i
				}
			}
		}
		return -1*/
}
