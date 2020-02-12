package string

import "strings"

// 58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")

	if len(words) < 1 {
		return 0
	}

	return len(words[len(words)-1])
}
