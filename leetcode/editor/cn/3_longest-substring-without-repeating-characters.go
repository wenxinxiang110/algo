package main

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 11267 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	length := 0
	hash := make(map[byte]bool)
	var left, right = 0, -1

	for ; left < len(s); left++ {
		if left > 0 {
			delete(hash, s[left-1])
		}
		for ; right+1 < len(s) && !hash[s[right+1]]; right++ {
			hash[s[right+1]] = true
		}
		length = max(length, right-left+1)
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
