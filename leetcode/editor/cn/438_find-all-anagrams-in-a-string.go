package main

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//
//
// 示例 1:
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 1856 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return []int{}
	}
	if len(s) < len(p) {
		return []int{}
	}
	var result []int

	// init p char map
	pBits := [26]uint8{}
	for _, pc := range p {
		pBits[pc-'a']++
	}

	sBits := [26]uint8{}
	// set int sBits
	var left, right = 0, len(p) - 1
	for i := left; i < right; i++ {
		sBits[s[i]-'a']++
	}
	for ; right < len(s); left++ {
		sBits[s[right]-'a']++
		if sBits == pBits {
			result = append(result, left)
		}
		sBits[s[left]-'a']--
		right++
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
