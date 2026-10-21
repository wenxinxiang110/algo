package main

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
//
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按 非递减顺序 排序
//
//
//
//
// 进阶：
//
//
// 请你设计时间复杂度为 O(n) 的算法解决本问题
//
//
// Related Topics 数组 双指针 排序 👍 1032 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

func sortedSquares(nums []int) []int {
	return sortedSquaresV2(nums)
}

func sortedSquaresV1(nums []int) []int {
	// 找到分界点，拆分开负数和正数
	firstPositiveIndex := -1
	for i, v := range nums {
		if v >= 0 {
			firstPositiveIndex = i
			break
		}
	}
	// 拆分为正负两个有序数组
	var negatives, positives []int
	if firstPositiveIndex == -1 {
		negatives = make([]int, len(nums))
	} else {
		negatives = make([]int, firstPositiveIndex)
		positives = make([]int, len(nums)-firstPositiveIndex)
	}
	for i := 0; i < len(negatives); i++ {
		negatives[len(negatives)-i-1] = nums[i] * nums[i]
	}
	for i := 0; i < len(positives); i++ {
		positives[i] = nums[i+firstPositiveIndex] * nums[i+firstPositiveIndex]
	}

	// 合并正负数组
	var merge = make([]int, 0)
	var neg, pos = 0, 0
	for neg < len(negatives) && pos < len(positives) {
		if negatives[neg] <= positives[pos] {
			merge = append(merge, negatives[neg])
			neg++
		} else {
			merge = append(merge, positives[pos])
			pos++
		}
	}
	if neg != len(negatives) {
		merge = append(merge, negatives[neg:]...)
	} else {
		merge = append(merge, positives[pos:]...)
	}

	// 合并两个有序数组
	return merge
}

// 双指针方法，从数组两端同时出发，这两个数的平方里有一个是最大的
func sortedSquaresV2(nums []int) []int {
	var merge = make([]int, len(nums))
	// 从 merge数组的最后一位出发
	var mergeIdx = len(merge) - 1
	var neg, pos = 0, len(nums) - 1
	for neg <= pos {
		negS := nums[neg] * nums[neg]
		posS := nums[pos] * nums[pos]
		if negS <= posS {
			merge[mergeIdx] = posS
			pos--
		} else {
			merge[mergeIdx] = negS
			neg++
		}
		mergeIdx--
	}
	return merge
}

//leetcode submit region end(Prohibit modification and deletion)
