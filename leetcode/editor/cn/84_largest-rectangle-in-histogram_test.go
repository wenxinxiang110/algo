package main

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	t.Run("示例1：[2,1,5,6,2,3]", func(t *testing.T) {
		heights := []int{2, 1, 5, 6, 2, 3}
		got := largestRectangleArea(heights)
		expected := 10

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("示例2：[2,4]", func(t *testing.T) {
		heights := []int{2, 4}
		got := largestRectangleArea(heights)
		expected := 4

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("单调递增：[1,2,3,4,5]", func(t *testing.T) {
		heights := []int{1, 2, 3, 4, 5}
		got := largestRectangleArea(heights)
		expected := 9 // 最大矩形：3*3=9

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("单调递减：[5,4,3,2,1]", func(t *testing.T) {
		heights := []int{5, 4, 3, 2, 1}
		got := largestRectangleArea(heights)
		expected := 9 // 最大矩形：3*3=9

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("全相同：[3,3,3,3]", func(t *testing.T) {
		heights := []int{3, 3, 3, 3}
		got := largestRectangleArea(heights)
		expected := 12 // 最大矩形：3*4=12

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("单元素：[5]", func(t *testing.T) {
		heights := []int{5}
		got := largestRectangleArea(heights)
		expected := 5 // 最大矩形：5*1=5

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("包含零值：[0,2,0,3,0]", func(t *testing.T) {
		heights := []int{0, 2, 0, 3, 0}
		got := largestRectangleArea(heights)
		expected := 3 // 最大矩形：3*1=3

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("全零：[0,0,0,0]", func(t *testing.T) {
		heights := []int{0, 0, 0, 0}
		got := largestRectangleArea(heights)
		expected := 0 // 最大矩形面积为0

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("空数组：[]", func(t *testing.T) {
		heights := []int{}
		got := largestRectangleArea(heights)
		expected := 0 // 空数组的最大矩形面积为0

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("峰值在中间：[1,3,2,4,1]", func(t *testing.T) {
		heights := []int{1, 3, 2, 4, 1}
		got := largestRectangleArea(heights)
		expected := 6 // 最大矩形：2*3=6

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("大数值：[10000,10000,10000]", func(t *testing.T) {
		heights := []int{10000, 10000, 10000}
		got := largestRectangleArea(heights)
		expected := 30000 // 最大矩形：10000*3=30000

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("边界最大值：[10000]", func(t *testing.T) {
		heights := []int{10000}
		got := largestRectangleArea(heights)
		expected := 10000 // 最大矩形：10000*1=10000

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("复杂情况：[6,7,5,2,4,5,9,3]", func(t *testing.T) {
		heights := []int{6, 7, 5, 2, 4, 5, 9, 3}
		got := largestRectangleArea(heights)
		expected := 16 // 最大矩形：4*4=16

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("最小长度：[5]", func(t *testing.T) {
		heights := []int{5}
		got := largestRectangleArea(heights)
		expected := 5 // 最小长度数组的最大矩形面积

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("验证时间复杂度：大数组", func(t *testing.T) {
		// 创建一个大数组测试性能
		heights := make([]int, 100000)
		for i := range heights {
			heights[i] = i%100 + 1 // 生成1-100的循环值
		}

		got := largestRectangleArea(heights)

		// 验证结果非负且合理
		if got < 0 {
			t.Errorf("largestRectangleArea(大数组) = %d，期望非负值", got)
		}

		if got > 100000*10000 { // 最大可能面积：长度*最大高度
			t.Errorf("largestRectangleArea(大数组) = %d，结果超出合理范围", got)
		}
	})

	t.Run("锯齿形：[1,10,1,10,1,10]", func(t *testing.T) {
		heights := []int{1, 10, 1, 10, 1, 10}
		got := largestRectangleArea(heights)
		expected := 10 // 最大矩形：10*1=10

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("阶梯形：[1,2,3,2,1]", func(t *testing.T) {
		heights := []int{1, 2, 3, 2, 1}
		got := largestRectangleArea(heights)
		expected := 6 // 最大矩形：2*3=6

		if got != expected {
			t.Errorf("largestRectangleArea(%v) = %d, 期望 %d", heights, got, expected)
		}
	})

	t.Run("验证约束条件：长度和高度范围", func(t *testing.T) {
		// 测试数组长度边界
		heights := make([]int, 100000)
		for i := range heights {
			heights[i] = 10000 // 最大允许高度
		}

		got := largestRectangleArea(heights)
		expected := 100000 * 10000 // 最大可能面积

		if got != expected {
			t.Errorf("largestRectangleArea(最大数组) = %d, 期望 %d", got, expected)
		}
	})
}
