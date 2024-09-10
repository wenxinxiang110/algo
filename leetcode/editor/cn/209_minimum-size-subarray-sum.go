package main

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
//长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
// 示例 1：
//
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
//
// 示例 2：
//
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//
//
// 示例 3：
//
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
//
//
// 进阶：
//
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
//
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 2222 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) (minLen int) {
	var swl = 0 // 滑动窗口的起始位置
	var sum = 0 // 滑动窗口的总和
	// 往右移动滑动窗口
	for swr := 0; swr < len(nums); swr++ {
		sum += nums[swr] // 记录总和
		for sum >= target {
			size := swr - swl + 1
			if minLen == 0 || size < minLen {
				minLen = size
			}
			sum -= nums[swl]
			swl++
		}
	}
	return minLen
}

//leetcode submit region end(Prohibit modification and deletion)

// 暴力解法 O(n^2)
func minSubArrayLenV1(target int, nums []int) int {
	return minSubArrayLenV2(target, nums)
}

// 滑动窗口（双指针解法）
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.md
func minSubArrayLenV2(target int, nums []int) (minLen int) {
	var swl = 0 // 滑动窗口的起始位置
	var sum = 0 // 滑动窗口的总和
	// 往右移动滑动窗口
	for swr := 0; swr < len(nums); swr++ {
		sum += nums[swr] // 记录总和
		for sum >= target {
			size := swr - swl + 1
			if minLen == 0 || size < minLen {
				minLen = size
			}
			sum -= nums[swl]
			swl++
		}
	}
	return minLen
}
