package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	t.Run("示例1：n=3", func(t *testing.T) {
		n := 3
		got := generateParenthesis(n)

		expected := []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}

		if len(got) != len(expected) {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), len(expected))
		}

		// 检查每个组合是否都在预期结果中
		for _, combination := range got {
			if !containsString(expected, combination) {
				t.Errorf("generateParenthesis(%d) 返回了意外的组合: %s", n, combination)
			}
		}

		// 检查是否所有预期组合都被返回
		for _, expectedCombination := range expected {
			if !containsString(got, expectedCombination) {
				t.Errorf("generateParenthesis(%d) 缺少组合: %s", n, expectedCombination)
			}
		}
	})

	t.Run("示例2：n=1", func(t *testing.T) {
		n := 1
		got := generateParenthesis(n)

		expected := []string{
			"()",
		}

		if len(got) != len(expected) {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("generateParenthesis(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("n=2", func(t *testing.T) {
		n := 2
		got := generateParenthesis(n)

		expected := []string{
			"(())",
			"()()",
		}

		if len(got) != len(expected) {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), len(expected))
		}

		for _, combination := range got {
			if !containsString(expected, combination) {
				t.Errorf("generateParenthesis(%d) 返回了意外的组合: %s", n, combination)
			}
		}

		for _, expectedCombination := range expected {
			if !containsString(got, expectedCombination) {
				t.Errorf("generateParenthesis(%d) 缺少组合: %s", n, expectedCombination)
			}
		}
	})

	t.Run("n=4", func(t *testing.T) {
		n := 4
		got := generateParenthesis(n)

		// n=4 应该有14个有效组合
		expectedCount := 14
		if len(got) != expectedCount {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), expectedCount)
		}

		// 检查所有组合是否都是唯一的
		if !allCombinationsUnique(got) {
			t.Errorf("generateParenthesis(%d) 返回了重复的组合", n)
		}

		// 检查每个组合是否都是有效的括号
		for _, combination := range got {
			if !isValidParentheses(combination) {
				t.Errorf("generateParenthesis(%d) 返回了无效的括号组合: %s", n, combination)
			}
		}
	})

	t.Run("n=0", func(t *testing.T) {
		n := 0
		got := generateParenthesis(n)

		// n=0 应该返回空数组
		expected := []string{}

		if len(got) != len(expected) {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), len(expected))
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("generateParenthesis(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("边界情况：n=8", func(t *testing.T) {
		n := 8
		got := generateParenthesis(n)

		// n=8 应该有1430个有效组合（卡特兰数 C_8 = 1430）
		expectedCount := 1430
		if len(got) != expectedCount {
			t.Errorf("generateParenthesis(%d) 返回 %d 个组合，期望 %d 个", n, len(got), expectedCount)
		}

		// 检查所有组合是否都是唯一的
		if !allCombinationsUnique(got) {
			t.Errorf("generateParenthesis(%d) 返回了重复的组合", n)
		}

		// 检查每个组合是否都是有效的括号
		for _, combination := range got {
			if !isValidParentheses(combination) {
				t.Errorf("generateParenthesis(%d) 返回了无效的括号组合: %s", n, combination)
			}
		}
	})
}

// containsString 检查字符串数组中是否包含指定的字符串
func containsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

// allCombinationsUnique 检查所有括号组合是否都是唯一的
//func allCombinationsUnique(combinations []string) bool {
//	seen := make(map[string]bool)
//	for _, combination := range combinations {
//		if seen[combination] {
//			return false
//		}
//		seen[combination] = true
//	}
//	return true
//}

// isValidParentheses 检查括号字符串是否有效
func isValidParentheses(s string) bool {
	balance := 0
	for _, ch := range s {
		if ch == '(' {
			balance++
		} else if ch == ')' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
