package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	t.Run("示例1：[1,2,3]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := subsets(nums)

		// 预期的所有子集
		expected := [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}

		// 使用改进的比较方法
		if !compareSubsetsIgnoringOrder(got, expected) {
			t.Errorf("subsets(%v) = %v, 期望包含所有子集 %v", nums, got, expected)
		}
	})

	t.Run("示例2：[0]", func(t *testing.T) {
		nums := []int{0}
		got := subsets(nums)

		expected := [][]int{
			{},
			{0},
		}

		if !compareSubsetsIgnoringOrder(got, expected) {
			t.Errorf("subsets(%v) = %v, 期望包含所有子集 %v", nums, got, expected)
		}
	})

	t.Run("空数组", func(t *testing.T) {
		nums := []int{}
		got := subsets(nums)

		// 空数组应该返回一个包含空数组的二维数组
		expected := [][]int{{}}

		if len(got) != len(expected) {
			t.Errorf("subsets(%v) 返回 %d 个子集，期望 %d 个", nums, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("subsets(%v) = %v, 期望 %v", nums, got, expected)
		}
	})

	t.Run("单个元素：[5]", func(t *testing.T) {
		nums := []int{5}
		got := subsets(nums)

		expected := [][]int{
			{},
			{5},
		}

		if !compareSubsetsIgnoringOrder(got, expected) {
			t.Errorf("subsets(%v) = %v, 期望包含所有子集 %v", nums, got, expected)
		}
	})

	t.Run("四个元素：[1,2,3,4]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		got := subsets(nums)

		// 4个元素应该有2^4 = 16个子集
		expectedCount := 16
		if len(got) != expectedCount {
			t.Errorf("subsets(%v) 返回 %d 个子集，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查所有子集是否都是唯一的
		if !allSubsetsUnique(got) {
			t.Errorf("subsets(%v) 返回了重复的子集", nums)
		}

		// 检查每个子集是否都是原数组的子集
		for _, subset := range got {
			if !isSubset(subset, nums) {
				t.Errorf("子集 %v 不是原数组 %v 的子集", subset, nums)
			}
		}

		// 检查是否包含了所有可能的子集
		if !containsAllSubsets(got, nums) {
			t.Errorf("subsets(%v) 缺少某些子集", nums)
		}
	})

	t.Run("包含负数：[0,-1,2]", func(t *testing.T) {
		nums := []int{0, -1, 2}
		got := subsets(nums)

		// 3个元素应该有2^3 = 8个子集
		expectedCount := 8
		if len(got) != expectedCount {
			t.Errorf("subsets(%v) 返回 %d 个子集，期望 %d 个", nums, len(got), expectedCount)
		}

		// 检查所有子集是否都是唯一的
		if !allSubsetsUnique(got) {
			t.Errorf("subsets(%v) 返回了重复的子集", nums)
		}

		// 检查每个子集是否都是原数组的子集
		for _, subset := range got {
			if !isSubset(subset, nums) {
				t.Errorf("子集 %v 不是原数组 %v 的子集", subset, nums)
			}
		}
	})

	t.Run("验证幂集大小", func(t *testing.T) {
		testCases := []struct {
			nums          []int
			expectedCount int
		}{
			{[]int{}, 1},            // 2^0 = 1
			{[]int{1}, 2},           // 2^1 = 2
			{[]int{1, 2}, 4},        // 2^2 = 4
			{[]int{1, 2, 3}, 8},     // 2^3 = 8
			{[]int{1, 2, 3, 4}, 16}, // 2^4 = 16
		}

		for _, tc := range testCases {
			got := subsets(tc.nums)
			if len(got) != tc.expectedCount {
				t.Errorf("subsets(%v) 返回 %d 个子集，期望 %d 个", tc.nums, len(got), tc.expectedCount)
			}
		}
	})
}

// compareSubsetsIgnoringOrder 比较两个子集集合是否相同，忽略顺序
func compareSubsetsIgnoringOrder(got, expected [][]int) bool {
	if len(got) != len(expected) {
		return false
	}

	// 对两个集合进行标准化排序
	normalizedGot := normalizeSubsets(got)
	normalizedExpected := normalizeSubsets(expected)

	// 比较标准化后的结果
	return reflect.DeepEqual(normalizedGot, normalizedExpected)
}

// normalizeSubsets 对子集集合进行标准化排序
func normalizeSubsets(subsets [][]int) [][]int {
	// 深拷贝
	result := make([][]int, len(subsets))
	for i, subset := range subsets {
		result[i] = make([]int, len(subset))
		copy(result[i], subset)
	}

	// 对每个子集排序
	for _, subset := range result {
		sort.Ints(subset)
	}

	// 对子集数组排序（按长度，然后按字典序）
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) != len(result[j]) {
			return len(result[i]) < len(result[j])
		}
		for k := 0; k < len(result[i]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})

	return result
}

// allSubsetsUnique 检查所有子集是否都是唯一的
func allSubsetsUnique(subsets [][]int) bool {
	seen := make(map[string]bool)
	for _, subset := range subsets {
		// 对子集排序后生成key
		sortedSubset := make([]int, len(subset))
		copy(sortedSubset, subset)
		sort.Ints(sortedSubset)

		key := ""
		for _, num := range sortedSubset {
			key += string(rune(num)) + ","
		}
		if seen[key] {
			return false
		}
		seen[key] = true
	}
	return true
}

// isSubset 检查subset是否是nums的子集
func isSubset(subset []int, nums []int) bool {
	elementSet := make(map[int]bool)
	for _, num := range nums {
		elementSet[num] = true
	}

	for _, num := range subset {
		if !elementSet[num] {
			return false
		}
	}
	return true
}

// containsAllSubsets 检查结果是否包含所有可能的子集
func containsAllSubsets(result [][]int, nums []int) bool {
	// 生成所有可能的子集
	allPossible := generateAllSubsets(nums)

	// 标准化两个集合
	normalizedResult := normalizeSubsets(result)
	normalizedExpected := normalizeSubsets(allPossible)

	// 比较标准化后的结果
	return reflect.DeepEqual(normalizedResult, normalizedExpected)
}

// generateAllSubsets 生成所有可能的子集（用于验证）
func generateAllSubsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(int, []int)

	backtrack = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, []int{})
	return result
}
