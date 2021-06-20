package string

import (
	"strings"
)

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//示例：
//
//s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
//
//提示：你可以假定该字符串只包含小写字母。

func firstUniqChar(s string) int {

	array := []rune(s)
	for i := 0; i < len(array); i++ {
		if strings.IndexRune(s, array[i]) == strings.LastIndexAny(s, string(array[i])) {
			return i
		}
	}

	return -1
}
