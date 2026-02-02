package main

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
// 示例 3：
//
//
//输入：nums = [1,0,1,2]
//输出：3
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 并查集 数组 哈希表 👍 2736 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	maxL := 0
	// key是以 num 为 key 的最长连续序列数
	numsMap := make(map[int]bool, len(nums))

	// 初始化为0
	for _, num := range nums {
		numsMap[num] = true
	}
	for num := range numsMap {
		// 说明不是序列的起点，先不用算
		if numsMap[num-1] {
			continue
		}
		next := num + 1
		for numsMap[next] {
			next++
		}
		maxL = max(maxL, next-num)
	}
	return maxL
}

//leetcode submit region end(Prohibit modification and deletion)
