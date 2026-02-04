package main

//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1], k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3], k = 3
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
// Related Topics 数组 哈希表 前缀和 👍 3016 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) (count int) {
	if len(nums) == 0 {
		return 0
	}
	pres := 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pres += nums[i]
		// 计算前面出现过的对数
		count += mp[pres-k]
		mp[pres]++
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)

// 暴力
func subarrayViolent(nums []int, k int) (count int) {
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return
}

// 进阶：前缀和
func subarrayPrefixSum(nums []int, k int) (count int) {
	if len(nums) == 0 {
		return 0
	}
	pres := make([]int, len(nums)+1)
	pres[0] = 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pres[i+1] = pres[i] + nums[i]
		// 计算前面出现过的对数
		count += mp[pres[i+1]-k]
		mp[pres[i+1]]++
	}

	return
}

// 进阶：前缀和,优化空间复杂度
func subarrayPrefixSumV2(nums []int, k int) (count int) {
	if len(nums) == 0 {
		return 0
	}
	pres := 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pres += nums[i]
		// 计算前面出现过的对数
		count += mp[pres-k]
		mp[pres]++
	}

	return
}
