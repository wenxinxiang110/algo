package recursion

//给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
//示例 1：
//
//输入：n = 1
//输出：true
//解释：20 = 1
//示例 2：
//
//输入：n = 16
//输出：true
//解释：24 = 16
//示例 3：
//
//输入：n = 3
//输出：false
//示例 4：
//
//输入：n = 4
//输出：true
//示例 5：
//
//输入：n = 5
//输出：false
//
//提示：
//
//-2^31 <= n <= 2^31 - 1

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// 其实就是统计2进制表示下，是否只有一位为1
	count := 0
	for n != 0 {
		count += n % 2
		if count > 1 {
			return false
		}
		n /= 2
	}
	return true
}
