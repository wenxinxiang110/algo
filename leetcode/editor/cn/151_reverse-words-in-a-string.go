package main

import (
	"slices"
	"strings"
)

//给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
//
//
//
// 示例 1：
//
//
//输入：s = "the sky is blue"
//输出："blue is sky the"
//
//
// 示例 2：
//
//
//输入：s = "  hello world  "
//输出："world hello"
//解释：反转后的字符串中不能存在前导空格和尾随空格。
//
//
// 示例 3：
//
//
//输入：s = "a good   example"
//输出："example good a"
//解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词
//
//
//
//
//
//
//
// 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
//
// Related Topics 双指针 字符串 👍 1216 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

//func reverse(b []byte) {
//	for i := 0; i < len(b)/2; i++ {
//		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
//	}
//}

func reverseWords(s string) string {
	b := []byte(s)

	// 去除多余空格
	b = removeExtraSpaces(b)

	// 整体旋转
	reverse(b)
	// 按单词旋转
	var start = 0
	for end := start + 1; end <= len(b); end++ {
		if end == len(b) || b[end] == ' ' {
			reverse(b[start:end])
			start = end + 1
			end = start
		}
	}

	return string(b)
}

// 定义: 前后空格、中间连续多个空格代表多余空格
func removeExtraSpaces(b []byte) []byte {
	slow := 0
	for i := 0; i < len(b); i++ {
		// 判断是否多余空格
		if b[i] != ' ' {
			// 如果是开头的单词则不需要手动补0,否则需要手动补
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			// 保存单词
			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	return b[:slow]
}

// leetcode submit region end(Prohibit modification and deletion)
func reverseWordsByStd(s string) string {
	// 拆分
	split := strings.Split(s, " ")
	// 反转
	slices.Reverse(split)
	// 过滤空的
	var cur int = 0
	for i := 0; i < len(split); i++ {
		if split[i] == "" {
			continue
		} else {
			if i != cur {
				split[cur] = split[i]
			}
			cur++
		}
	}
	// 重组
	var buf strings.Builder
	for i := 0; i < cur; i++ {
		buf.WriteString(split[i])
		if i != cur-1 {
			buf.WriteByte(' ')
		}
	}

	return buf.String()
}

func mySplitBySpace(s []byte) (after [][]byte) {
	const sep = ' '
	var start = 0
	for start < len(s)-1 {
		for start != sep {
			start++
		}
		var end = start + 1
		for end < len(s) && s[end] != sep {
			end++
		}
		if end <= len(s) {
			after = append(after, s[start:end])
			start = end
			continue
		} else {
			after = append(after, s[start:len(s)])
			break
		}
	}
	return
}

func removeSpaces(b []byte) []byte {
	slow := 0
	for i := 0; i < len(b); i++ {
		// 判断是否多余空格
		if b[i] != ' ' {
			if i != slow {
				b[slow] = b[i]
			}
			slow++
		}
	}
	return b[:slow]
}
