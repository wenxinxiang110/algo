package main

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
//如果 needle 不是 haystack 的一部分，则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
//
//
// 示例 2：
//
//
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack 和 needle 仅由小写英文字符组成
//
//
// Related Topics 双指针 字符串 字符串匹配 👍 2298 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	b := []byte(haystack)
	n := []byte(needle)
	next := Next(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && b[i] != n[j+1] {
			j = next[j]
		}
		if b[i] == n[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - j
		}

	}
	return -1
}

// Next kmp前缀表,所有值-1
func Next(s string) []int {
	b := []byte(s)
	n := len(s)
	table := make([]int, n)

	//j指向前缀末尾位置，同时表示该位置上的最长前缀后缀长度
	var j = -1
	// 第一个位置可以省略查找的过程
	table[0] = j

	//i指向后缀末尾位置
	for i := 1; i < n; i++ {
		for j >= 0 && b[i] != b[j+1] {
			j = table[j]
		}
		if b[i] == b[j+1] { // 找到相同的前后缀
			j++
		}
		table[i] = j // 将j（前缀的长度）赋给next[i]
	}

	return table
}

//leetcode submit region end(Prohibit modification and deletion)
