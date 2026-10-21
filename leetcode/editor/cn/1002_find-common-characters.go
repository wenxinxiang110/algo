package main

//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案
//。
//
//
//
// 示例 1：
//
//
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
//
//
// 示例 2：
//
//
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] 由小写英文字母组成
//
//
// Related Topics 数组 哈希表 字符串 👍 376 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func commonChars(words []string) []string {
	return commonCharsV2(words)
}

func commonCharsV1(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	var table = make([][26]int, len(words))

	// set first
	// 第一行也可以用来记录整体的字符出现的最小次数
	for _, r := range []rune(words[0]) {
		table[0][r-'a']++
	}
	// set next
	for i := 1; i < len(words); i++ {
		for _, r := range []rune(words[i]) {
			table[i][r-'a']++
		}
		for j := 0; j < 26; j++ {
			if table[i][j] < table[0][j] {
				table[0][j] = table[i][j]
			}
		}
	}

	var result []string
	for i := 0; i < 26; i++ {
		for ; table[0][i] > 0; table[0][i]-- {
			result = append(result, string(rune('a'+i)))
		}
	}

	return result
}

// v1基础上优化空间复杂度为 O(1)
func commonCharsV2(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	var first = [26]int{}

	// set first
	// 第一行也可以用来记录整体的字符出现的最小次数
	for _, r := range words[0] {
		first[r-'a']++
	}

	var other = [26]int{}
	// set next
	for i := 1; i < len(words); i++ {
		for _, r := range []rune(words[i]) {
			other[r-'a']++
		}
		for j := 0; j < 26; j++ {
			first[j] = min(first[j], other[j])
			other[j] = 0
		}
	}

	var result []string
	for i := 0; i < 26; i++ {
		for ; first[i] > 0; first[i]-- {
			result = append(result, string(rune('a'+i)))
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
