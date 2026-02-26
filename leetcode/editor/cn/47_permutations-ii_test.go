package main

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	t.Run("示例1：[1,1,2]", func(t *testing.T) {
		nums := []int{1, 1, 2}
		got := permuteUnique(nums)

		// 预期的所有不重复排列
		expected := [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}

		if len(got) != len(expected) {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		// 检查每个排列是否都在预期结果中
		for _, perm := range got {
			if !containsPermutation(expected, perm) {
				t.Errorf("permuteUnique(%v) 返回了意外的排列: %v", nums, perm)
			}
		}

		// 检查是否所有预期排列都被返回
		for _, expectedPerm := range expected {
			if !containsPermutation(got, expectedPerm) {
				t.Errorf("permuteUnique(%v) 缺少排列: %v", nums, expectedPerm)
			}
		}

		// 检查是否没有重复排列
		if !allPermutationsUnique(got) {
			t.Errorf("permuteUnique(%v) 返回了重复的排列", nums)
		}
	})

	t.Run("示例2：[1,2,3]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := permuteUnique(nums)

		// 3个不同元素应该有3! = 6个排列
		expectedCount := 6
		if len(got) != expectedCount {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查所有排列是否都是唯一的
		if !allPermutationsUnique(got) {
			t.Errorf("permuteUnique(%v) 返回了重复的排列", nums)
		}

		// 检查每个排列是否包含所有原始元素
		for _, perm := range got {
			if !containsAllElements(perm, nums) {
				t.Errorf("排列 %v 不包含所有原始元素 %v", perm, nums)
			}
		}
	})

	t.Run("多个重复元素：[2,2,1,1]", func(t *testing.T) {
		nums := []int{2, 2, 1, 1}
		got := permuteUnique(nums)

		// 4个元素，其中2个1和2个2，应该有6个不重复排列
		// 排列数 = 4! / (2! * 2!) = 24 / 4 = 6
		expectedCount := 6
		if len(got) != expectedCount {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查是否没有重复排列
		if !allPermutationsUnique(got) {
			t.Errorf("permuteUnique(%v) 返回了重复的排列", nums)
		}

		// 检查每个排列是否包含所有原始元素
		for _, perm := range got {
			if !containsAllElements(perm, nums) {
				t.Errorf("排列 %v 不包含所有原始元素 %v", perm, nums)
			}
		}
	})

	t.Run("所有元素相同：[1,1,1]", func(t *testing.T) {
		nums := []int{1, 1, 1}
		got := permuteUnique(nums)

		// 所有元素相同，应该只有1个排列
		expected := [][]int{
			{1, 1, 1},
		}

		if len(got) != len(expected) {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("permuteUnique(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("单个元素：[5]", func(t *testing.T) {
		nums := []int{5}
		got := permuteUnique(nums)

		expected := [][]int{
			{5},
		}

		if len(got) != len(expected) {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("permuteUnique(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("空数组", func(t *testing.T) {
		nums := []int{}
		got := permuteUnique(nums)

		// 空数组应该返回一个包含空数组的二维数组
		expected := [][]int{{}}

		if len(got) != len(expected) {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("permuteUnique(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("包含负数和零：[0,-1,0]", func(t *testing.T) {
		nums := []int{0, -1, 0}
		got := permuteUnique(nums)

		// 3个元素，其中2个0和1个-1，应该有3个不重复排列
		// 排列数 = 3! / 2! = 6 / 2 = 3
		expectedCount := 3
		if len(got) != expectedCount {
			t.Errorf("permuteUnique(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查是否没有重复排列
		if !allPermutationsUnique(got) {
			t.Errorf("permuteUnique(%v) 返回了重复的排列", nums)
		}

		// 检查每个排列是否包含所有原始元素
		for _, perm := range got {
			if !containsAllElements(perm, nums) {
				t.Errorf("排列 %v 不包含所有原始元素 %v", perm, nums)
			}
		}
	})
}
