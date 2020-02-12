package string

import "strings"

//14. 最长公共前缀
func longestCommonPrefix(strs []string) (com string) {

	// 特殊处理
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	com = LCP(strs[0], strs[1:])

	return
}

// 查找两个字符串pre和post的最长公共前缀
func FindCom(pre, post string) (result string) {

	// 特殊处理
	if pre == "" || post == "" {
		return ""
	}

	for i := range post {
		if strings.Index(pre, post[:i+1]) != 0 {
			break
		}
		result = post[:i+1]
	}

	return
}

// 时间复杂度O(mn)
func LCP(this string, others []string) (result string) {

	// 特殊处理
	if this == "" || len(others) == 0 {
		return this
	}

	// 先找前两个字符串的最长公共前缀
	next := FindCom(this, others[0])

	if next != "" {
		result = LCP(next, others[1:])
	}

	return
}
