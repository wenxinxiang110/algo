package main

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	t.Run("示例1：n=4", func(t *testing.T) {
		n := 4
		got := solveNQueens(n)
		expected := [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("solveNQueens(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("示例2：n=1", func(t *testing.T) {
		n := 1
		got := solveNQueens(n)
		expected := [][]string{{"Q"}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("solveNQueens(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("n=2（无解）", func(t *testing.T) {
		n := 2
		got := solveNQueens(n)
		expected := [][]string{}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("solveNQueens(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("n=3（无解）", func(t *testing.T) {
		n := 3
		got := solveNQueens(n)
		expected := [][]string{}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("solveNQueens(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("n=5（有解）", func(t *testing.T) {
		n := 5
		got := solveNQueens(n)

		// 验证结果格式和数量
		if len(got) == 0 {
			t.Errorf("solveNQueens(%d) 应该返回解，但返回了空数组", n)
		}

		// 验证每个解的格式
		for i, solution := range got {
			if len(solution) != n {
				t.Errorf("solveNQueens(%d) 解[%d]的长度应该为%d，实际为%d", n, i, n, len(solution))
			}

			for j, row := range solution {
				if len(row) != n {
					t.Errorf("solveNQueens(%d) 解[%d]行[%d]的长度应该为%d，实际为%d", n, i, j, n, len(row))
				}

				// 验证每行只有一个Q
				qCount := 0
				for _, char := range row {
					if char == 'Q' {
						qCount++
					}
				}
				if qCount != 1 {
					t.Errorf("solveNQueens(%d) 解[%d]行[%d]应该恰好有一个Q，实际有%d个", n, i, j, qCount)
				}
			}
		}
	})

	t.Run("n=6（有解）", func(t *testing.T) {
		n := 6
		got := solveNQueens(n)

		// 验证至少有一个解
		if len(got) == 0 {
			t.Errorf("solveNQueens(%d) 应该返回解，但返回了空数组", n)
		}

		// 验证解的格式正确
		validateSolutionFormat(t, n, got)
	})

	t.Run("n=7（有解）", func(t *testing.T) {
		n := 7
		got := solveNQueens(n)

		if len(got) == 0 {
			t.Errorf("solveNQueens(%d) 应该返回解，但返回了空数组", n)
		}

		validateSolutionFormat(t, n, got)
	})

	t.Run("n=8（经典八皇后）", func(t *testing.T) {
		n := 8
		got := solveNQueens(n)

		// 八皇后应该有92个解
		expectedCount := 92
		if len(got) != expectedCount {
			t.Errorf("solveNQueens(%d) 应该有%d个解，实际有%d个", n, expectedCount, len(got))
		}

		validateSolutionFormat(t, n, got)
	})

	t.Run("n=9（最大边界）", func(t *testing.T) {
		n := 9
		got := solveNQueens(n)

		if len(got) == 0 {
			t.Errorf("solveNQueens(%d) 应该返回解，但返回了空数组", n)
		}

		validateSolutionFormat(t, n, got)
	})

	t.Run("验证解的合法性", func(t *testing.T) {
		n := 4
		got := solveNQueens(n)

		for i, solution := range got {
			// 验证皇后不互相攻击
			if !isValidSolution(solution) {
				t.Errorf("solveNQueens(%d) 解[%d]不合法：皇后互相攻击", n, i)
			}
		}
	})

	t.Run("边界最小值：n=1", func(t *testing.T) {
		n := 1
		got := solveNQueens(n)
		expected := [][]string{{"Q"}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("solveNQueens(%d) = %v, 期望 %v", n, got, expected)
		}
	})

	t.Run("边界最大值：n=9", func(t *testing.T) {
		n := 9
		got := solveNQueens(n)

		// 验证解的格式
		validateSolutionFormat(t, n, got)

		// 验证至少有一个解
		if len(got) == 0 {
			t.Errorf("solveNQueens(%d) 应该返回解，但返回了空数组", n)
		}
	})

	t.Run("验证所有解的独特性", func(t *testing.T) {
		n := 4
		got := solveNQueens(n)

		// 检查是否有重复解
		seen := make(map[string]bool)
		for i, solution := range got {
			solutionKey := ""
			for _, row := range solution {
				solutionKey += row + "|"
			}

			if seen[solutionKey] {
				t.Errorf("solveNQueens(%d) 解[%d]与其他解重复", n, i)
			}
			seen[solutionKey] = true
		}
	})
}

// validateSolutionFormat 验证解的格式是否正确
func validateSolutionFormat(t *testing.T, n int, solutions [][]string) {
	t.Helper()

	for i, solution := range solutions {
		if len(solution) != n {
			t.Errorf("solveNQueens(%d) 解[%d]的长度应该为%d，实际为%d", n, i, n, len(solution))
		}

		for j, row := range solution {
			if len(row) != n {
				t.Errorf("solveNQueens(%d) 解[%d]行[%d]的长度应该为%d，实际为%d", n, i, j, n, len(row))
			}

			// 验证每行字符只能是Q或.
			for _, char := range row {
				if char != 'Q' && char != '.' {
					t.Errorf("solveNQueens(%d) 解[%d]行[%d]包含非法字符：%c", n, i, j, char)
				}
			}

			// 验证每行只有一个Q
			qCount := 0
			for _, char := range row {
				if char == 'Q' {
					qCount++
				}
			}
			if qCount != 1 {
				t.Errorf("solveNQueens(%d) 解[%d]行[%d]应该恰好有一个Q，实际有%d个", n, i, j, qCount)
			}
		}
	}
}

// isValidSolution 验证解是否合法（皇后不互相攻击）
func isValidSolution(solution []string) bool {
	n := len(solution)

	// 检查行和列
	rows := make([]int, n)
	cols := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution[i][j] == 'Q' {
				rows[i]++
				cols[j]++

				if rows[i] > 1 || cols[j] > 1 {
					return false
				}
			}
		}
	}

	// 检查对角线
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution[i][j] == 'Q' {
				// 检查左上到右下对角线
				for k := 1; i+k < n && j+k < n; k++ {
					if solution[i+k][j+k] == 'Q' {
						return false
					}
				}

				// 检查右上到左下对角线
				for k := 1; i+k < n && j-k >= 0; k++ {
					if solution[i+k][j-k] == 'Q' {
						return false
					}
				}
			}
		}
	}

	return true
}
