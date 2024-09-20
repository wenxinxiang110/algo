package main

//给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 935 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{}, len(nums1))
	for _, n := range nums1 {
		set1[n] = struct{}{}
	}
	set2 := make(map[int]struct{}, len(nums2))
	for _, n := range nums2 {
		set2[n] = struct{}{}
	}
	var result []int
	for n := range set1 {
		_, ok := set2[n]
		if ok {
			result = append(result, n)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
