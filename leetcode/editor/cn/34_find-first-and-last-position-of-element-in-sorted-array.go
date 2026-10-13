package main

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 3216 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	//分解为两个问题：
	// 找到第一个等于 target 的位置
	// 找到最后一个大于target 的位置 -1

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 找到第一个位置
	var start = binarySearchFirst(nums, target)
	if start == -1 {
		return []int{-1, -1}
	}

	var end = binarySearchLast(nums, target)
	return []int{start, end}

}

// 二分查找有序数组中最后一个等于 target的元素
func binarySearchLast(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		// 这个加1 避免死循环 ？ 为什么
		// 当 left = 0, right = 1 时，mid = (0 + 1) / 2 = 0，left 无法更新
		// 加1 后 mid = (0 + 1 + 1) / 2 = 1，left 可以更新
		mid := (left + right + 1) / 2
		// 1. 即使找到 target, 大于 target的肯定在右边
		if nums[mid] == target {
			left = mid
		} else if nums[mid] > target {
			// 可以排除右边界，target在左边和 mid里
			right = mid - 1
		} else {
			//if nums[mid] < target  target肯定在右边
			left = mid
		}
	}
	return left
}

// 二分查找有序数组中最后一个等于 target的元素
//func binarySearchLast(nums []int, target int) int {
//	var left, right = 0, len(nums) - 1
//	for left < right {
//		mid := (left + right + 1) / 2 // 使用上取整，避免死循环
//		if nums[mid] == target {
//			left = mid // 向右收缩，寻找最后一个等于target的位置
//		} else if nums[mid] > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	// 检查是否找到target
//	if left >= 0 && left < len(nums) && nums[left] == target {
//		return left
//	}
//	return -1
//}

// 二分查找有序数组中第一个等于 target的元素
func binarySearchFirst(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		// 1. 即使找到 target,左边可能也有等于 target的
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			// 可以拍出右边界
			right = mid - 1
		} else {
			//if nums[mid] < target  target肯定在右边
			left = mid + 1
		}
	}
	// 也有可能没找到
	if nums[left] == target {
		return left
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
