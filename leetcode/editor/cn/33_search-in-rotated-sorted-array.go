package main

//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，使数组变为 [nums[k],
//nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1
//,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 3331 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	return searchRotateK(nums, target)
}

func searchRotateK(nums []int, target int) int {

	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		// 已经找到了
		if nums[mid] == target {
			return mid
		}
		//没找到，需要从左或者右边的有序序列里搜索

		// 左边有序
		if nums[left] <= nums[mid] {
			//且要搜索的值在左边
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1

			}
		}
	}
	return -1
}

// 二分查找，但是找 k的时候已经 O(n)了
func searchRotateKBad(nums []int, target int) int {
	//	 find k
	var k int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			k = i + 1
			break
		}
	}

	// 辅助函数，在一个向右旋转k 的数组中查找
	var numGet = func(idx int) int {
		return nums[(idx+k)%len(nums)]
	}

	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if numGet(mid) < target {
			left = mid + 1
		} else if numGet(mid) > target {
			right = mid - 1
		} else {
			return (mid + k) % len(nums)
		}

	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
