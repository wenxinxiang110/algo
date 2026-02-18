package main

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。 请你实现时间复杂度为
//O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//解释：范围 [1,2] 中的数字都在数组中。
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//解释：1 在数组中，但 2 没有。
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//解释：最小的正数 1 没有出现。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
// Related Topics 数组 哈希表 👍 2506 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	// O1空间的关键，是复用入参中的数组空间
	// 基本思路: 在长度为 N 的数组中，没有出现过的最小正整数，应该在 [1, N+1] 范围内
	// 那么，我们就可以利用数组本身的空间，来存储这个范围
	// 比如，遍历到 1，那么 nums[0] = 1
	// 遍历到 2，那么 nums[1] = 2
	// ....
	// 遍历到 N，那么 nums[N-1] = N
	// 遍历完成后，再遍历一次数组，如果 i != nums[i-1]，那么 i 就是答案
	// 第一：先把非正数和 0 移除
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	//  第二：把正数放到对应的位置
	// 不能直接把对应位置的数置为 0或者清楚
	for i := 0; i < len(nums); i++ {
		num := abs(nums[i])
		if num <= len(nums) {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

//leetcode submit region end(Prohibit modification and deletion)
