package main

import (
	"math"
)

//给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，
//返回空字符串 ""。
//
// 测试用例保证答案唯一。
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//O(m + n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 3482 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) (ans string) {

	var cntS, cntT = make(map[byte]int), make(map[byte]int)
	//
	for _, tc := range []byte(t) {
		cntT[tc] += 1
	}
	// 判断是否覆盖
	isCovered := func() bool {
		for tc, tCount := range cntT {
			if cntS[tc] < tCount {
				return false
			}
		}
		return true
	}
	minLength := math.MaxInt16

	var left, right = 0, 0
	for ; right < len(s); right++ {
		cntS[s[right]] += 1

		for isCovered() && left <= right {
			if right-left+1 < minLength {
				minLength = min(minLength, right-left+1)
				ans = s[left : right+1]
			}
			if _, ok := cntS[s[left]]; ok {
				cntS[s[left]] -= 1
			}
			left++
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
