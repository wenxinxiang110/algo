package array

//给定两个数组，编写一个函数来计算它们的交集。
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//示例 2：
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//说明：
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	m := make(map[int]int)
	for _, i := range nums1 {
		v := m[i]
		if v == 0 {
			m[i] = 1
		}
	}
	for _, i := range nums2 {
		v := m[i]
		if v == 1 {
			m[i] = 2
		}
	}
	var res []int
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
