package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	t.Run("示例1：[2,3,6,7], target=7", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		got := combinationSum(candidates, target)

		expected := [][]int{
			{2, 2, 3},
			{7},
		}

		if !compareCombinationsIgnoringOrder(got, expected) {
			t.Errorf("combinationSum(%v, %d) = %v, 期望 %v", candidates, target, got, expected)
		}
	})

	t.Run("示例2：[2,3,5], target=8", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		got := combinationSum(candidates, target)

		expected := [][]int{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		}

		if !compareCombinationsIgnoringOrder(got, expected) {
			t.Errorf("combinationSum(%v, %d) = %v, 期望 %v", candidates, target, got, expected)
		}
	})

	t.Run("示例3：[2], target=1", func(t *testing.T) {
		candidates := []int{2}
		target := 1
		got := combinationSum(candidates, target)

		// 无法组成和为1的组合，应该返回空切片或nil
		// 使用len(got) == 0来检查是否为空结果
		if len(got) != 0 {
			t.Errorf("combinationSum(%v, %d) = %v, 期望空结果", candidates, target, got)
		}
	})

	t.Run("空数组", func(t *testing.T) {
		candidates := []int{}
		target := 5
		got := combinationSum(candidates, target)

		// 空数组应该返回空结果
		if len(got) != 0 {
			t.Errorf("combinationSum(%v, %d) = %v, 期望空结果", candidates, target, got)
		}
	})

	t.Run("目标为0", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 0
		got := combinationSum(candidates, target)

		// 目标为0时，应该返回包含空组合的切片
		expected := [][]int{{}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("combinationSum(%v, %d) = %v, 期望 %v", candidates, target, got, expected)
		}
	})

	t.Run("单个元素等于目标：[3], target=3", func(t *testing.T) {
		candidates := []int{3}
		target := 3
		got := combinationSum(candidates, target)

		expected := [][]int{{3}}

		if !compareCombinationsIgnoringOrder(got, expected) {
			t.Errorf("combinationSum(%v, %d) = %v, 期望 %v", candidates, target, got, expected)
		}
	})

	t.Run("多个组合：[2,3,5], target=7", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 7
		got := combinationSum(candidates, target)

		// 可能的组合：{2,2,3}, {2,5}
		expected := [][]int{
			{2, 2, 3},
			{2, 5},
		}

		if !compareCombinationsIgnoringOrder(got, expected) {
			t.Errorf("combinationSum(%v, %d) = %v, 期望 %v", candidates, target, got, expected)
		}
	})

	t.Run("验证组合唯一性：[2,3,6,7], target=7", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		got := combinationSum(candidates, 7)

		if !allIntCombinationsUnique(got) {
			t.Errorf("combinationSum(%v, 7) 返回了重复的组合: %v", candidates, got)
		}
	})

	t.Run("验证组合正确性：[2,3,5], target=8", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		got := combinationSum(candidates, 8)

		// 验证每个组合的和都等于target
		for _, combination := range got {
			sum := 0
			for _, num := range combination {
				sum += num
			}
			if sum != 8 {
				t.Errorf("组合 %v 的和为 %d，不等于目标值 8", combination, sum)
			}
		}
	})

	t.Run("较大目标值：[2,3,5], target=10", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 10
		got := combinationSum(candidates, target)

		// 验证所有组合的和都等于10
		for _, combination := range got {
			sum := 0
			for _, num := range combination {
				sum += num
			}
			if sum != target {
				t.Errorf("组合 %v 的和为 %d，不等于目标值 %d", combination, sum, target)
			}
		}

		// 验证组合唯一性
		if !allIntCombinationsUnique(got) {
			t.Errorf("combinationSum(%v, %d) 返回了重复的组合", candidates, target)
		}
	})

	t.Run("包含较大数字：[1,2,5,10], target=15", func(t *testing.T) {
		candidates := []int{1, 2, 5, 10}
		target := 15
		got := combinationSum(candidates, target)

		// 验证所有组合的和都等于15
		for _, combination := range got {
			sum := 0
			for _, num := range combination {
				sum += num
			}
			if sum != target {
				t.Errorf("组合 %v 的和为 %d，不等于目标值 %d", combination, sum, target)
			}
		}
	})
}

// compareCombinationsIgnoringOrder 比较两个组合切片是否相同，忽略顺序
func compareCombinationsIgnoringOrder(got, expected [][]int) bool {
	if len(got) != len(expected) {
		return false
	}

	// 对每个组合进行排序，然后对整个结果进行排序
	sortedGot := sortCombinations(got)
	sortedExpected := sortCombinations(expected)

	// 比较排序后的结果
	return reflect.DeepEqual(sortedGot, sortedExpected)
}

// sortCombinations 对组合切片进行排序
func sortCombinations(combinations [][]int) [][]int {
	// 先复制切片
	sorted := make([][]int, len(combinations))
	for i, combination := range combinations {
		sorted[i] = make([]int, len(combination))
		copy(sorted[i], combination)
		// 对每个组合进行排序
		sort.Ints(sorted[i])
	}

	// 对整个结果按字典序排序
	sort.Slice(sorted, func(i, j int) bool {
		for k := 0; k < len(sorted[i]) && k < len(sorted[j]); k++ {
			if sorted[i][k] != sorted[j][k] {
				return sorted[i][k] < sorted[j][k]
			}
		}
		return len(sorted[i]) < len(sorted[j])
	})

	return sorted
}

// allIntCombinationsUnique 检查所有组合是否都是唯一的
func allIntCombinationsUnique(combinations [][]int) bool {
	seen := make(map[string]bool)

	for _, combination := range combinations {
		// 对组合进行排序以确保唯一性比较
		sortedCombination := make([]int, len(combination))
		copy(sortedCombination, combination)
		sort.Ints(sortedCombination)

		// 将排序后的组合转换为字符串作为键
		key := ""
		for _, num := range sortedCombination {
			key += string(rune(num)) + ","
		}

		if seen[key] {
			return false
		}
		seen[key] = true
	}
	return true
}

// validateCombinationSum 验证组合是否满足条件
func validateCombinationSum(combinations [][]int, candidates []int, target int) bool {
	for _, combination := range combinations {
		sum := 0
		for _, num := range combination {
			// 检查数字是否在候选数组中
			found := false
			for _, candidate := range candidates {
				if num == candidate {
					found = true
					break
				}
			}
			if !found {
				return false
			}
			sum += num
		}
		if sum != target {
			return false
		}
	}
	return true
}
