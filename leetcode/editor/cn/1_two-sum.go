package main

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
//
//
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
//
// Related Topics 数组 哈希表 👍 18856 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	//m := make(map[int]int, len(nums))
	//for idx, num := range nums {
	//	// find target
	//	nextIdx, ok := m[target-num]
	//	// has been found, return
	//	if ok {
	//		return []int{nextIdx, idx}
	//	}
	//	// not found, se into map
	//	m[num] = idx
	//}
	return twoSumHot(nums, target)
}
func twoSumHot(nums []int, target int) []int {
	// 暴力
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {func twoSumHot(nums []int, target int) []int {
	//	// 暴力
	//	for i := 0; i < len(nums)-1; i++ {
	//		for j := i + 1; j < len(nums); j++ {
	//			if nums[i]+nums[j] == target {
	//				return []int{i, j}
	//			}
	//		}
	//	}
	//	return nil
	//}
	//			return []int{i, j}
	//		}
	//	}
	//}

	// hash
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if twins, ok := m[target-nums[i]]; ok {
			return []int{twins, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func twoSumV1(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 在生成 map的时候就可以查找数据了
func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, num := range nums {
		// find target
		nextIdx, ok := m[target-num]
		// has been found, return
		if ok {
			return []int{nextIdx, idx}
		}
		// not found, se into map
		m[num] = idx
	}
	return nil
}
