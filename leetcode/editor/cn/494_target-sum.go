package main

import (
	"math"
)

//给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//
//
// 示例 2：
//
//
//输入：nums = [1], target = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
//
//
// Related Topics 数组 动态规划 回溯 👍 2033 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < int(math.Abs(float64(target))) {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	pkg := (sum + target) / 2

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, pkg+1)
	}
	// 初始化行
	if nums[0] <= pkg {
		dp[0][nums[0]] = 1
	}
	// 初始化列
	zeroNums := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNums++
		}
		dp[i][0] = int(math.Pow(2, float64(zeroNums)))
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= pkg; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i] {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(nums)-1][pkg]
}

//leetcode submit region end(Prohibit modification and deletion)
