package main

//编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
//
//
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 19
//输出：true
//解释：
//1² + 9² = 82
//8² + 2² = 68
//6² + 8² = 100
//1² + 0² + 0² = 1
//
//
// 示例 2：
//
//
//输入：n = 2
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
//
// Related Topics 哈希表 数学 双指针 👍 1607 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	var record = make(map[int]struct{})
	for {
		var sum = getSum(n)
		if sum == 1 {
			return true
		}
		_, ok := record[sum]
		if ok {
			return false
		} else {
			record[sum] = struct{}{}
		}
		n = sum
	}
}

func getSum(n int) (sum int) {
	iterDigest(n, func(digest int) {
		sum += digest * digest
	})
	return
}

//leetcode submit region end(Prohibit modification and deletion)

// 取数值各个位上的单数
func iterDigest(n int, do func(digest int)) {
	for n != 0 {
		do(n % 10)
		n /= 10
	}
}
