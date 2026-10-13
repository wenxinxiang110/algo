package main

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	t.Run("示例1：[1,2,3]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := permute(nums)

		// 预期的所有排列
		expected := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}

		if len(got) != len(expected) {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		// 检查每个排列是否都在预期结果中
		for _, perm := range got {
			if !containsPermutation(expected, perm) {
				t.Errorf("permute(%v) 返回了意外的排列: %v", nums, perm)
			}
		}

		// 检查是否所有预期排列都被返回
		for _, expectedPerm := range expected {
			if !containsPermutation(got, expectedPerm) {
				t.Errorf("permute(%v) 缺少排列: %v", nums, expectedPerm)
			}
		}
	})

	t.Run("示例2：[0,1]", func(t *testing.T) {
		nums := []int{0, 1}
		got := permute(nums)

		expected := [][]int{
			{0, 1},
			{1, 0},
		}

		if len(got) != len(expected) {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		for _, perm := range got {
			if !containsPermutation(expected, perm) {
				t.Errorf("permute(%v) 返回了意外的排列: %v", nums, perm)
			}
		}

		for _, expectedPerm := range expected {
			if !containsPermutation(got, expectedPerm) {
				t.Errorf("permute(%v) 缺少排列: %v", nums, expectedPerm)
			}
		}
	})

	t.Run("示例3：[1]", func(t *testing.T) {
		nums := []int{1}
		got := permute(nums)

		expected := [][]int{
			{1},
		}

		if len(got) != len(expected) {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("permute(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("四个元素的排列", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		got := permute(nums)

		// 4个元素应该有4! = 24个排列
		expectedCount := 24
		if len(got) != expectedCount {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查所有排列是否都是唯一的
		if !allPermutationsUnique(got) {
			t.Errorf("permute(%v) 返回了重复的排列", nums)
			t.Errorf("%v", got)
		}

		// 检查每个排列是否包含所有原始元素
		for _, perm := range got {
			if !containsAllElements(perm, nums) {
				t.Errorf("排列 %v 不包含所有原始元素 %v", perm, nums)
			}
		}
	})

	t.Run("空数组", func(t *testing.T) {
		nums := []int{}
		got := permute(nums)

		// 空数组应该返回一个包含空数组的二维数组
		expected := [][]int{{}}

		if len(got) != len(expected) {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("permute(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("负数元素", func(t *testing.T) {
		nums := []int{-1, 0, 1}
		got := permute(nums)

		// 3个元素应该有3! = 6个排列
		expectedCount := 6
		if len(got) != expectedCount {
			t.Errorf("permute(%v) 返回 %d 个排列，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查所有排列是否都是唯一的
		if !allPermutationsUnique(got) {
			t.Errorf("permute(%v) 返回了重复的排列", nums)
		}

		// 检查每个排列是否包含所有原始元素
		for _, perm := range got {
			if !containsAllElements(perm, nums) {
				t.Errorf("排列 %v 不包含所有原始元素 %v", perm, nums)
			}
		}
	})
}

// containsPermutation 检查二维数组中是否包含指定的排列
func containsPermutation(permutations [][]int, target []int) bool {
	for _, perm := range permutations {
		if reflect.DeepEqual(perm, target) {
			return true
		}
	}
	return false
}

// allPermutationsUnique 检查所有排列是否都是唯一的
func allPermutationsUnique(permutations [][]int) bool {
	seen := make(map[string]bool)
	for _, perm := range permutations {
		key := ""
		for _, num := range perm {
			key += string(rune(num)) + ","
		}
		if seen[key] {
			return false
		}
		seen[key] = true
	}
	return true
}

// containsAllElements 检查排列是否包含原始数组的所有元素
func containsAllElements(perm []int, original []int) bool {
	if len(perm) != len(original) {
		return false
	}

	elementCount := make(map[int]int)
	for _, num := range original {
		elementCount[num]++
	}

	for _, num := range perm {
		elementCount[num]--
		if elementCount[num] < 0 {
			return false
		}
	}

	return true
}
