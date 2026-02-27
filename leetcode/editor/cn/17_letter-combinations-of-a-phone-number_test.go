package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Run("示例1：\"23\"", func(t *testing.T) {
		digits := "23"
		got := letterCombinations(digits)

		expected := []string{
			"ad", "ae", "af",
			"bd", "be", "bf",
			"cd", "ce", "cf",
		}

		if !compareStringSlicesIgnoringOrder(got, expected) {
			t.Errorf("letterCombinations(\"%s\") = %v, 期望 %v", digits, got, expected)
		}
	})

	t.Run("示例2：\"2\"", func(t *testing.T) {
		digits := "2"
		got := letterCombinations(digits)

		expected := []string{"a", "b", "c"}

		if !compareStringSlicesIgnoringOrder(got, expected) {
			t.Errorf("letterCombinations(\"%s\") = %v, 期望 %v", digits, got, expected)
		}
	})

	t.Run("空字符串", func(t *testing.T) {
		digits := ""
		got := letterCombinations(digits)

		// 空字符串应该返回空切片
		var expected []string = nil

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("letterCombinations(\"%s\") = %v, 期望 %v", digits, got, expected)
		}
	})

	t.Run("单个数字：\"7\"", func(t *testing.T) {
		digits := "7"
		got := letterCombinations(digits)

		// 数字7对应4个字母
		expected := []string{"p", "q", "r", "s"}

		if !compareStringSlicesIgnoringOrder(got, expected) {
			t.Errorf("letterCombinations(\"%s\") = %v, 期望 %v", digits, got, expected)
		}
	})

	t.Run("两个数字：\"79\"", func(t *testing.T) {
		digits := "79"
		got := letterCombinations(digits)

		// 7对应4个字母，9对应4个字母，总共16种组合
		expectedCount := 4 * 4
		if len(got) != expectedCount {
			t.Errorf("letterCombinations(\"%s\") 返回 %d 个组合，期望 %d 个", digits, len(got), expectedCount)
		}

		// 检查所有组合是否唯一
		if !allCombinationsUnique(got) {
			t.Errorf("letterCombinations(\"%s\") 返回了重复的组合", digits)
		}
	})

	t.Run("三个数字：\"234\"", func(t *testing.T) {
		digits := "234"
		got := letterCombinations(digits)

		// 2对应3个字母，3对应3个字母，4对应3个字母，总共27种组合
		expectedCount := 3 * 3 * 3
		if len(got) != expectedCount {
			t.Errorf("letterCombinations(\"%s\") 返回 %d 个组合，期望 %d 个", digits, len(got), expectedCount)
		}

		// 检查所有组合是否唯一
		if !allCombinationsUnique(got) {
			t.Errorf("letterCombinations(\"%s\") 返回了重复的组合", digits)
		}
	})

	t.Run("四个数字：\"5678\"", func(t *testing.T) {
		digits := "5678"
		got := letterCombinations(digits)

		// 5对应3个字母，6对应3个字母，7对应4个字母，8对应3个字母，总共108种组合
		expectedCount := 3 * 3 * 4 * 3
		if len(got) != expectedCount {
			t.Errorf("letterCombinations(\"%s\") 返回 %d 个组合，期望 %d 个", digits, len(got), expectedCount)
		}

		// 检查所有组合是否唯一
		if !allCombinationsUnique(got) {
			t.Errorf("letterCombinations(\"%s\") 返回了重复的组合", digits)
		}
	})

	t.Run("验证组合数量", func(t *testing.T) {
		testCases := []struct {
			digits        string
			expectedCount int
		}{
			{"", 0},       // 空字符串
			{"2", 3},      // 3个字母
			{"23", 9},     // 3*3=9
			{"79", 16},    // 4*4=16
			{"234", 27},   // 3*3*3=27
			{"5678", 108}, // 3*3*4*3=108
		}

		for _, tc := range testCases {
			got := letterCombinations(tc.digits)
			if len(got) != tc.expectedCount {
				t.Errorf("letterCombinations(\"%s\") 返回 %d 个组合，期望 %d 个", tc.digits, len(got), tc.expectedCount)
			}
		}
	})

	t.Run("验证组合内容：\"23\"", func(t *testing.T) {
		digits := "23"
		got := letterCombinations(digits)

		// 验证所有组合都包含正确的字母
		expectedLetters := map[string]bool{
			"ad": true, "ae": true, "af": true,
			"bd": true, "be": true, "bf": true,
			"cd": true, "ce": true, "cf": true,
		}

		for _, combination := range got {
			if !expectedLetters[combination] {
				t.Errorf("letterCombinations(\"%s\") 返回了无效的组合: %s", digits, combination)
			}
		}
	})
}

// compareStringSlicesIgnoringOrder 比较两个字符串切片是否相同，忽略顺序
func compareStringSlicesIgnoringOrder(got, expected []string) bool {
	if len(got) != len(expected) {
		return false
	}

	// 对两个切片进行排序
	sortedGot := make([]string, len(got))
	copy(sortedGot, got)
	sort.Strings(sortedGot)

	sortedExpected := make([]string, len(expected))
	copy(sortedExpected, expected)
	sort.Strings(sortedExpected)

	// 比较排序后的结果
	return reflect.DeepEqual(sortedGot, sortedExpected)
}

// allCombinationsUnique 检查所有组合是否都是唯一的
func allCombinationsUnique(combinations []string) bool {
	seen := make(map[string]bool)
	for _, combination := range combinations {
		if seen[combination] {
			return false
		}
		seen[combination] = true
	}
	return true
}

// generateExpectedCombinations 生成预期的所有组合（用于验证）
func generateExpectedCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var result []string
	var backtrack func(int, string)

	backtrack = func(index int, current string) {
		if index == len(digits) {
			result = append(result, current)
			return
		}

		digit := digits[index]
		letters := phoneMap[digit]
		for _, letter := range letters {
			backtrack(index+1, current+letter)
		}
	}

	backtrack(0, "")
	return result
}
