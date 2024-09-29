package main

//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
//
//
// 示例 2：
//
//
//输入：s = "abcd", k = 2
//输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
//
//
// Related Topics 双指针 字符串 👍 617 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	l := len(s)
	b := []byte(s)
	for i := 0; i < l; i += 2 * k {
		if l-i < k {
			reverse(b[i:l])
		} else {
			reverse(b[i : i+k])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func reverseStrV1(s string, k int) string {
	l := len(s)
	idx := 0
	runes := []rune(s)
	after := make([]rune, len(runes))
	for i := 0; i < l; i += 2 * k {
		if l-i < k {
			for j := l - 1; j >= i; j-- {
				after[idx] = runes[j]
				idx++
			}
		} else {
			for j := 0; j < k; j++ {
				after[idx] = runes[i+k-j-1]
				idx++
			}
			// 原样复制
			for j := 0; j < k && i+k+j < l; j++ {
				after[idx] = runes[i+k+j]
				idx++
			}
		}
	}
	return string(after)
}

func reverseStrV2(s string, k int) string {
	l := len(s)
	runes := []rune(s)
	after := make([]rune, 0)
	for i := 0; i < l; i += 2 * k {
		if l-i < k {
			for j := l - 1; j >= i; j-- {
				after = append(after, runes[j])
			}
		} else {
			for j := 0; j < k; j++ {
				after = append(after, runes[i+k-j-1])
			}
			// 原样复制
			end := i + 2*k
			if end > l {
				end = l
			}
			after = append(after, runes[i+k:end]...)
		}
	}
	return string(after)
}

func reverseStrV3(s string, k int) string {
	l := len(s)
	b := []byte(s)
	for i := 0; i < l; i += 2 * k {
		if l-i < k {
			reverse(b[i:l])
		} else {
			reverse(b[i : i+k])
		}
	}
	return string(b)
}
