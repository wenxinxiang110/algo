package main

import (
	"testing"
)

func TestSearchRotateK(t *testing.T) {
	t.Run("示例1：target存在", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		got := searchRotateK(nums, target)
		expected := 4

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("示例2：target不存在", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("示例3：单元素不存在", func(t *testing.T) {
		nums := []int{1}
		target := 0
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("单元素存在", func(t *testing.T) {
		nums := []int{1}
		target := 1
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("未旋转数组：target存在", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
		target := 4
		got := searchRotateK(nums, target)
		expected := 4

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("未旋转数组：target不存在", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
		target := 8
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("旋转点在最左：target存在", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
		target := 0
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("旋转点在最右：target存在", func(t *testing.T) {
		nums := []int{7, 0, 1, 2, 3, 4, 5, 6}
		target := 7
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("旋转点在中间：target存在", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
		target := 7
		got := searchRotateK(nums, target)
		expected := 3

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("target在旋转点左侧", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
		target := 5
		got := searchRotateK(nums, target)
		expected := 1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("target在旋转点右侧", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
		target := 1
		got := searchRotateK(nums, target)
		expected := 5

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("边界最小值：数组长度=1", func(t *testing.T) {
		nums := []int{-10000}
		target := -10000
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("边界最大值：数组长度=5000", func(t *testing.T) {
		// 创建长度为5000的旋转数组
		nums := make([]int, 5000)
		for i := 0; i < 5000; i++ {
			nums[i] = i
		}
		// 旋转数组，旋转点在2500
		rotated := make([]int, 5000)
		copy(rotated[0:2500], nums[2500:])
		copy(rotated[2500:], nums[0:2500])

		target := 3750
		got := searchRotateK(rotated, target)
		expected := 1250 // 3750在旋转后的位置

		if got != expected {
			t.Errorf("searchRotateK(长度5000数组, %d) = %d, 期望 %d", target, got, expected)
		}
	})

	t.Run("负数值存在", func(t *testing.T) {
		nums := []int{-10, -5, -3, -1, 0, 2, 4, 6}
		target := -5
		got := searchRotateK(nums, target)
		expected := 1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("负数值不存在", func(t *testing.T) {
		nums := []int{-10, -5, -3, -1, 0, 2, 4, 6}
		target := -7
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("target小于数组最小值", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
		target := -1
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("target大于数组最大值", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
		target := 8
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("两元素数组：target存在", func(t *testing.T) {
		nums := []int{1, 0}
		target := 1
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("两元素数组：target不存在", func(t *testing.T) {
		nums := []int{1, 0}
		target := 2
		got := searchRotateK(nums, target)
		expected := -1

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("三元素数组：target在旋转点", func(t *testing.T) {
		nums := []int{2, 0, 1}
		target := 2
		got := searchRotateK(nums, target)
		expected := 0

		if got != expected {
			t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, target, got, expected)
		}
	})

	t.Run("验证数组旋转属性", func(t *testing.T) {
		testCases := []struct {
			nums     []int
			expected bool
		}{
			{[]int{4, 5, 6, 7, 0, 1, 2}, true},
			{[]int{0, 1, 2, 3, 4, 5, 6}, true},
			{[]int{6, 0, 1, 2, 3, 4, 5}, true},
			{[]int{1}, true},
			{[]int{1, 0}, true},
		}

		for _, tc := range testCases {
			// 验证数组是否满足旋转排序数组的属性
			if !isRotatedSortedArray(tc.nums) {
				t.Errorf("数组%v不满足旋转排序数组属性", tc.nums)
			}
		}
	})

	t.Run("验证时间复杂度要求", func(t *testing.T) {
		// 创建大数组测试性能（虽然不是严格的O(log n)验证，但可以观察执行时间）
		nums := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			nums[i] = i
		}
		// 旋转数组
		rotated := make([]int, 1000)
		copy(rotated[0:500], nums[500:])
		copy(rotated[500:], nums[0:500])

		target := 750
		got := searchRotateK(rotated, target)
		expected := 250

		if got != expected {
			t.Errorf("searchRotateK(长度1000数组, %d) = %d, 期望 %d", target, got, expected)
		}
	})

	t.Run("验证函数正确性：所有元素都能找到", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2, 3}

		for i, expected := range nums {
			got := searchRotateK(nums, expected)
			if got != i {
				t.Errorf("searchRotateK(%v, %d) = %d, 期望 %d", nums, expected, got, i)
			}
		}
	})
}

// isRotatedSortedArray 验证数组是否满足旋转排序数组的属性
func isRotatedSortedArray(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	// 找到旋转点
	rotatePoint := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			rotatePoint = i + 1
			break
		}
	}

	// 验证旋转点左侧是递增的
	for i := 0; i < rotatePoint-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}

	// 验证旋转点右侧是递增的
	for i := rotatePoint; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}

	// 验证旋转点右侧的最大值小于等于旋转点左侧的最小值
	if rotatePoint > 0 && rotatePoint < len(nums) {
		if nums[len(nums)-1] > nums[0] {
			return false
		}
	}

	return true
}
