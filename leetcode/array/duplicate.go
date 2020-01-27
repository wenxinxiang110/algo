package array

//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
//双指针，思路类似上面的数组去重
func removeDuplicates(nums []int) int {
	i := 1
	//m := make(map[int]int)
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			//m[nums[j]] = 1
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

//给定一个整数数组，判断是否存在重复元素。
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			return true
		}
	}
	return false
}

//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
//链接：https://leetcode-cn.com/problems/contains-duplicate-ii
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, ok := m[v]; !ok {
			m[v] = i
		} else {
			if i-index < k {
				return true
			}
			m[v] = i
		}
	}
	return false
}
