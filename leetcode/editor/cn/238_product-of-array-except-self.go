package main

//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//
//
// 示例 2:
//
//
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 输入 保证 数组 answer[i] 在 32 位 整数范围内
//
//
//
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
//
// Related Topics 数组 前缀和 👍 2186 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	return productExceptSelfSpaceO1(nums)
}

// 不能用除法，则可以前缀乘积数组乘以后最乘积数组
func productExceptSelfLR(nums []int) (ans []int) {
	var L, R = make([]int, len(nums)), make([]int, len(nums))

	// init left
	L = append(L, 1)
	for i := 0; i < len(nums); i++ {
		L[i] = L[i-1] * nums[i-1]
	}

	// init right
	R[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, L[i]*R[i])
	}

	return ans
}

// 基于上面的优化，O1额外存储空间
func productExceptSelfSpaceO1(nums []int) (ans []int) {
	ans = make([]int, len(nums))
	// init left
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// init right
	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R *= nums[i]
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
