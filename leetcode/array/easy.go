package array

//给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
//水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
//反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。
//链接：https://leetcode-cn.com/problems/flipping-an-image
func flipAndInvertImage(A [][]int) [][]int {
	//暴力 by me update by 题解：要交换的两个元素不相同，则反转之后又变一样了，不处理
	for i := 0; i < len(A); i++ {
		length := len(A[i])
		for j := 0; j < (length+1)/2; j++ {
			if A[i][length-1-j] == A[i][j] {
				A[i][j] ^= 1
				if j != length-1-j {
					A[i][length-1-j] ^= 1
				}
			}
		}
	}
	return A
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//二分查找 by me,有很细节问题，参考：https://leetcode-cn.com/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/
func SearchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	var mid int
	for start < end {
		mid = (start + end) >> 1
		if target < nums[mid] {
			end = mid - 1
		} else if target == nums[mid] {
			return mid
		} else {
			start = mid + 1
		}
	}
	if target <= nums[start] {
		return start
	} else {
		return start + 1
	}
}

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//链接：https://leetcode-cn.com/problems/plus-one
func PlusOne(digits []int) []int {
	loc, up := len(digits)-1, 1
	for ; up != 0 && loc >= 0; loc-- {
		up = (digits[loc] + 1) / 10
		digits[loc] = (digits[loc] + 1) % 10
	}
	if loc < 0 && up != 0 {
		//digits = append([]int{1}, digits...)
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}
	return digits
}

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
func Generate(numRows int) [][]int {
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		r[i] = make([]int, i+1)
		r[i][0] = 1
		r[i][i] = 1
		if i == 0 || i == 1 {
			continue
		}
		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}
	return r
}

//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//todo:unfinished
func getRow(rowIndex int) []int {
	return make([]int, rowIndex)
}

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
func twoSum(numbers []int, target int) []int {
	r := make([]int, 2)
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			r[0] = i + 1
			r[1] = j + 1
			return r
		} else if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		}
	}
	return r
}

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在众数。
func MajorityElement(nums []int) int {
	candidate := nums[0]
	count := 0
	for _, v := range nums {
		if v == candidate {
			count++
			continue
		}
		//
		if count > 0 {
			count--
		} else {
			candidate = v
			count = 1
		}

	}
	return candidate
}

//给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
//异或
func missingNumber(nums []int) int {
	miss := len(nums)
	for i := 0; i < len(nums); i++ {
		miss ^= i
		miss ^= nums[i]
	}
	return miss
}

//给定一个二进制数组， 计算其中最大连续1的个数。
//示例 1:
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
//链接：https://leetcode-cn.com/problems/max-consecutive-ones
func findMaxConsecutiveOnes(nums []int) int {
	length, maxLength := 0, 0
	for _, v := range nums {
		if v == 0 {
			length = 0
		} else {
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

//给定一个未经排序的整数数组，找到最长且连续的的递增序列。
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	length, maxLength := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			length = 1
		} else {
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

//示例 1：
//
//输入：2
//输出：1
//解释：F(2) = F(1) + F(0) = 1 + 0 = 1.
//示例 2：
//
//输入：3
//输出：2
//解释：F(3) = F(2) + F(1) = 1 + 1 = 2.
//示例 3：
//
//输入：4
//输出：3
//解释：F(4) = F(3) + F(2) = 2 + 1 = 3.
func fib(N int) int {
	if N <= 0 {
		return 0
	}
	if N < 3 {
		return 1
	}
	n1, n2 := 1, 1
	now := 0
	for i := 2; i < N; i++ {
		now = n1 + n2
		n1 = n2
		n2 = now
	}
	return now
}

//给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
//我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
//如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
//示例 1:
//输入:
//nums = [1, 7, 3, 6, 5, 6]
//输出: 3
//解释:
//索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
//同时, 3 也是第一个符合要求的中心索引。
//示例 2:
//输入:
//nums = [1, 2, 3]
//输出: -1
//解释:
//数组中不存在满足此条件的中心索引。
//说明:
//nums 的长度范围为 [0, 10000]。
//任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
//链接：https://leetcode-cn.com/problems/find-pivot-index
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		if left == sum-left-nums[i] {
			return i
		}
		left += nums[i]
	}
	return -1
}

//给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
//我的做法：冒泡排序
func thirdMax(nums []int) int {
	m := make(map[int]int, 3)
	var tmp int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
		m[nums[len(nums)-i-1]] = 1
		if len(m) == 3 {
			return nums[i]
		}
	}
	m[nums[0]] = 1
	if len(m) == 3 {
		return nums[0]
	}
	return nums[len(nums)-1]
}
