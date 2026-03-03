package main

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
//
//
// Related Topics 栈 数组 单调栈 👍 3078 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	return largestRectangleAreaFocus(heights)
}

// 暴力解法
func largestRectangleAreaFocus(height []int) (maxArea int) {
	for cur, curHeight := range height {
		//以 height[i]为起点，向左右找到宽度
		left := cur
		for left > 0 && height[left-1] >= curHeight {
			left--
		}
		right := cur
		for right < len(height)-1 && height[right+1] > curHeight {
			right++
		}

		area := curHeight * (right - left + 1)
		maxArea = max(maxArea, area)
	}
	return
}

// 单调栈解法
func largestRectangleAreaStack(height []int) (maxArea int) {
	if len(height) == 0 {
		return
	}
	if len(height) == 1 {
		return height[0]
	}

	var stack = []int{0}

	for i := 0; i < len(height); i++ {
		// 比较栈顶下标对应的
		if height[stack[len(stack)-1]] > height[i] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			area := height[pop] * (i - pop - 1)
			maxArea = max(maxArea, area)
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
