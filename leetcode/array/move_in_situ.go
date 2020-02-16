package array

import (
	"math"
)

//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//链接：https://leetcode-cn.com/problems/remove-element
//暴力 n^2 by me
// 双指针方法1： （这种可以保证顺序，但复制过程比较多）
// 当 nums[j]与给定的值相等时，递增 j 以跳过该元素。只要 nums[j] != val
// 我们就复制 nums[j]到 nums[i]并同时递增两个索引。重复这一过程，直到 j 到达数组的末尾，该数组的新长度为 i。
// 双指针法2： 当nums[i]==val时，交换nums和数组最后一个元素， 好处是复制比较少，缺点是无序
func RemoveElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//链接：https://leetcode-cn.com/problems/merge-sorted-array
//题解：因为nums1后面是空的，所以可以从后面开始扫描，这样就不用额外开一个CopyNums1了
func Merge(nums1 []int, m int, nums2 []int, n int) {
	//指向内存中同一个数组的切片，nums1改了copyNums1也改了,要用内置的copy函数
	//copyNums1 := make([]int, m)
	//copy(copyNums1, nums1[:m])
	//i, j := 0, 0
	//for ; i < m && j < n; {
	//	if copyNums1[i] < nums2[j] {
	//		nums1[i+j] = copyNums1[i]
	//		i++
	//	} else {
	//		nums1[i+j] = nums2[j]
	//		j++
	//	}
	//}
	//if i < m {
	//	copy(nums1[i+j:], copyNums1[i:])
	//}
	//if j < n {
	//	copy(nums1[i+j:], nums2[j:])
	//}
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode/
//两种方法：环状替换；先反转整个数组，再分别反转前面k个数字和后面n-k个数字
func Rotate(nums []int, k int) {
	L := len(nums)
	k = k % L
	for i := 0; i < L/2; i++ {
		nums[i], nums[L-i-1] = nums[L-i-1], nums[i]
	}
	reverse(nums, 0, k)
	reverse(nums, k, L)
	//nums = append(nums[k+1:], nums[:k+1]...)
	//fmt.Println(nums)
}
func reverse(nums []int, start, end int) {
	L := end - start
	for i := 0; i < L/2; i++ {
		nums[i+start], nums[end-i-1] = nums[end-i-1], nums[i+start]
	}
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

//给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//链接：https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[int(math.Abs(float64(nums[i]))-1)] > 0 {
			nums[int(math.Abs(float64(nums[i]))-1)] = 0 - nums[int(math.Abs(float64(nums[i]))-1)]
		}
	}
	r := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			r = append(r, i+1)
		}
	}
	return r
}

//在一排座位（ seats）中，1 代表有人坐在座位上，0 代表座位上是空的。
//至少有一个空座位，且至少有一人坐在座位上。
//亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。
//返回他到离他最近的人的最大距离。
//示例 1：
//输入：[1,0,0,0,1,0,1]
//输出：2
//示例 2：
//输入：[1,0,0,0]
//输出：3
//链接：https://leetcode-cn.com/problems/maximize-distance-to-closest-person
func maxDistToClosest(seats []int) int {
	max := 0
	start, end := 0, 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			end++
		} else {
			if start == 0 {
				max = end - start
			} else {
				max = int(math.Max(float64(max), float64((end-start)/2)))
			}
			start = i
		}
	}
	return max
}
