package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。
//
// 示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
//
// Related Topics 数组 二分查找 👍 1619 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/*func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		}

	}
	return -1
}*/

//leetcode submit region end(Prohibit modification and deletion)

// 二分查找重要的是区间，左闭右闭即[left, right]，或者左闭右开即[left, right)
func binarySearch(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	// 左闭右闭区间 [left, right]，循环条件 left <= right
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
