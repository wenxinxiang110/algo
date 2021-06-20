package string

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
	return firstUniqCharDoubleRange(s)
}

// 思路1: 第一次遍历，储存出现次数；第二次遍历，发现是只出现一次的，直接返回索引
func firstUniqCharDoubleRange(s string) int {
	// unique标记每个字符初始出现时候的索引
	var unique = [26]int{}
	array := []rune(s)

	for i := 0; i < len(array); i++ {
		unique[array[i]-'a']++
	}

	for i := 0; i < len(array); i++ {
		if unique[array[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
