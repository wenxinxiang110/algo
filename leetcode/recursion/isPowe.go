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
	// 其实就是统计2进制表示下，是否只有一位为1
	// n&(n-1)会移除最后一位1
	// n&(-n)会保留最后一位1
	return n > 0 && (n&(n-1)) == 0
	//return n > 0 && (n&(-n)) == n
}

//给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。
func isPowerOfThree(n int) bool {
	return IsPowerOf(n, 3)
}

// IsPowerOf 判断n是否为power的整数次幂
func IsPowerOf(n int, power int) bool {
	if n <= 0 {
		return false
	}
	for n%power == 0 {
		n /= power
	}
	return n == 1
}
