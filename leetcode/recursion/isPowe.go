package recursion

import (
	"math"
)

// IsPowerOf 判断n是否为power的整数次幂, 通用解法
func IsPowerOf(n int, power int) bool {
	if n <= 0 {
		return false
	}
	for n%power == 0 {
		n /= power
	}
	return n == 1
}

// isPowerOfLog 通用计算方法2: 通过换底公式推导,要处理float64精度问题，暂时没搞好
func isPowerOfLog(n, power int) bool {
	return math.Mod(math.Log(float64(n))/math.Log(float64(power))+math.SmallestNonzeroFloat64, 1) <= 2*math.SmallestNonzeroFloat64
}

// IsPowerOfByPrimeNumber n的最大值为math.MaxInt32, power为质数,通过计算power在这个领域的最大值来计算是否合适
func IsPowerOfByPrimeNumber(n, power int) bool {
	return n > 0 && int(math.Pow(float64(power), math.Round(math.Log(math.MaxInt32)/math.Log(float64(power)))))%n == 0
}

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
	return isPowerOfLog(n, 3)
}

// 给定一个整数，写一个函数来判断它是否是 4 的幂次方
func isPowerOfFour(n int) bool {
	// 在判断n是2的幂的基础上，再判断二进制表示里，唯一的一位1是否在偶数位上
	return n > 0 && n&(n-1) == 0 && n&0x55555555 != 0
}
