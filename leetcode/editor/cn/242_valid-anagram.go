package main

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。
//
//
//
// 示例 1:
//
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s 和 t 仅包含小写字母
//
//
//
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
// Related Topics 哈希表 字符串 排序 👍 946 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := [26]int{}
	for _, r := range s {
		sm[r-'a'] += 1
	}
	for _, r := range []rune(t) {
		sm[r-'a'] -= 1
	}
	for _, v := range sm {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := [26]int{}
	for _, r := range s {
		sm[r-'a'] += 1
	}
	tm := [26]int{}
	for _, r := range []rune(t) {
		tm[r-'a'] += 1
	}
	for i, v := range sm {
		if v != tm[i] {
			return false
		}
	}
	return true
}

func isAnagramV1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[rune]int, len([]rune(s)))
	for _, r := range []rune(s) {
		sm[r] = sm[r] + 1
	}
	tm := make(map[rune]int, len([]rune(t)))
	for _, r := range []rune(t) {
		tm[r] = tm[r] + 1
	}
	if len(sm) != len(tm) {
		return false
	}
	for k, v := range sm {
		if v != tm[k] {
			return false
		}
	}
	return true
}
