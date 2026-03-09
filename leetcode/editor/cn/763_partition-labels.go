package main

//给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab",
//"cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
//
// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//
// 返回一个表示每个字符串片段的长度的列表。
//
//
//示例 1：
//
//
//输入：s = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
//
// 示例 2：
//
//
//输入：s = "eccbbbbdec"
//输出：[10]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 仅由小写英文字母组成
//
//
// Related Topics 贪心 哈希表 双指针 字符串 👍 1393 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) (ans []int) {
	// charIndex[i]记录每个字符的起始和结束位置
	charIndex := [26][2]int{}
	for i, c := range s {
		// 记录每个字符的起始位置
		if charIndex[c-'a'][1] == 0 {
			charIndex[c-'a'][0] = i
		}
		// 记录每个字符的结束位置
		charIndex[c-'a'][1] = i
	}
	// 某个区间的开始和结束
	var start, end = 0, 0
	for i := 0; i < len(s); i++ {
		start = min(start, charIndex[s[i]-'a'][0])

		end = max(end, charIndex[s[i]-'a'][1])

		// 已经到分片的末尾了,记录分片大小，重置区间
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
