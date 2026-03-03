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
	return largestRectangleAreaStack(heights)
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

	var stack = []int{}
	empty := func() bool {
		return len(stack) == 0
	}
	push := func(i int) {
		stack = append(stack, i)
	}
	pop := func() int {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}
	top := func() int {
		return stack[len(stack)-1]
	}

	for i := 0; i < len(height); i++ {
		// 比较栈顶下标对应的，如果栈顶的位置的矩形高度比现在的高，说明他已经无法延伸，已经找到右边界了
		for !empty() && height[top()] > height[i] {
			//
			length := height[pop()]

			width := i
			if !empty() {
				// top此时是左边界
				width = i - top() - 1
			}
			maxArea = max(maxArea, length*width)
		}
		// 入栈
		push(i)
	}

	// 剩下的元素继续处理
	for !empty() {
		length := height[pop()]
		width := len(height)
		if !empty() {
			width = width - top() - 1
		}
		maxArea = max(maxArea, length*width)
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)
