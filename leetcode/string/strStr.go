package string

import "strings"

//28. 实现 strStr()
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
func strStr(haystack string, needle string) int {

	// 特殊处理
	if haystack == "" {
		return -1
	}
	if needle == "" {
		return 0
	}
	strings.Index("", "")

	// todo: unfinish
	return 0
}
