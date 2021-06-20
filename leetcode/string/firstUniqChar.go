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
	// unique标记每个字符初始出现时候的索引
	var unique = [26]int{}
	// unique的初始化,避免但值为0时，无法区分该key不存在和索引为0
	defaultVal := -1
	for i := 0; i < len(unique); i++ {
		unique[i] = defaultVal
	}

	array := []rune(s)
	for i := 0; i < len(array); i++ {
		uniqueIdx := array[i] - 'a'
		// 不重复，标记为该字符初次出现时的索引
		if unique[uniqueIdx] == defaultVal {
			unique[uniqueIdx] = i
		} else {
			// 不为0，代表已经为标记过了, 标记为重复
			unique[uniqueIdx] = defaultVal - 1
		}
	}
	// 查找不重复(unique[i]>0),且最早出现出现的字符索引
	var latest = defaultVal
	for i := 0; i < len(unique); i++ {
		if unique[i] < 0 {
			continue
		}
		if latest == defaultVal || unique[i] < latest {
			latest = unique[i]
		}
	}
	return latest
}
