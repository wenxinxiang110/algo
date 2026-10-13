package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	t.Run("示例1：目标值存在且有重复", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		got := searchRange(nums, target)
		expected := []int{3, 4}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("示例2：目标值不存在", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		got := searchRange(nums, target)
		expected := []int{-1, -1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("示例3：空数组", func(t *testing.T) {
		nums := []int{}
		target := 0
		got := searchRange(nums, target)
		expected := []int{-1, -1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("单元素数组：目标值存在", func(t *testing.T) {
		nums := []int{1}
		target := 1
		got := searchRange(nums, target)
		expected := []int{0, 0}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("单元素数组：目标值不存在", func(t *testing.T) {
		nums := []int{1}
		target := 2
		got := searchRange(nums, target)
		expected := []int{-1, -1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值出现在数组开头", func(t *testing.T) {
		nums := []int{1, 1, 2, 3, 4, 5}
		target := 1
		got := searchRange(nums, target)
		expected := []int{0, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值出现在数组结尾", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 5}
		target := 5
		got := searchRange(nums, target)
		expected := []int{4, 5}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值只出现一次", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		got := searchRange(nums, target)
		expected := []int{2, 2}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值多次连续出现", func(t *testing.T) {
		nums := []int{1, 2, 2, 2, 2, 3}
		target := 2
		got := searchRange(nums, target)
		expected := []int{1, 4}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值小于数组最小值", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 0
		got := searchRange(nums, target)
		expected := []int{-1, -1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("目标值大于数组最大值", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 6
		got := searchRange(nums, target)
		expected := []int{-1, -1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("负数值测试", func(t *testing.T) {
		nums := []int{-5, -3, -3, -1, 0, 2}
		target := -3
		got := searchRange(nums, target)
		expected := []int{1, 2}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, target, got, expected)
		}
	})

	t.Run("大数组边界测试", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			nums[i] = i
		}
		target := 500
		got := searchRange(nums, target)
		expected := []int{500, 500}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("searchRange(长度1000数组, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("验证函数正确性：所有唯一元素", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		for i, expectedValue := range nums {
			expected := []int{i, i}
			got := searchRange(nums, expectedValue)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, expectedValue, got, expected)
			}
		}
	})

	t.Run("验证函数正确性：重复元素", func(t *testing.T) {
		nums := []int{1, 2, 2, 3, 3, 3, 4}

		testCases := []struct {
			target   int
			expected []int
		}{
			{1, []int{0, 0}},
			{2, []int{1, 2}},
			{3, []int{3, 5}},
			{4, []int{6, 6}},
		}

		for _, tc := range testCases {
			got := searchRange(nums, tc.target)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("searchRange(%v, %d) = %v, 期望 %v", nums, tc.target, got, tc.expected)
			}
		}
	})
}
