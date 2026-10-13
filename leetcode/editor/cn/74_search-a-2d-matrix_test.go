package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	t.Run("示例1：target存在", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 3
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("示例2：target不存在", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 13
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("target在第一行第一个", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 1
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("target在最后一行最后一个", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 60
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("target小于矩阵最小值", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 0
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("target大于矩阵最大值", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 100
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("单行矩阵", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
		}
		target := 5
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("单列矩阵", func(t *testing.T) {
		matrix := [][]int{
			{1},
			{10},
			{23},
		}
		target := 10
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("1x1矩阵存在", func(t *testing.T) {
		matrix := [][]int{
			{5},
		}
		target := 5
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("1x1矩阵不存在", func(t *testing.T) {
		matrix := [][]int{
			{5},
		}
		target := 3
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("边界最小值：m=1,n=1", func(t *testing.T) {
		matrix := [][]int{
			{-10000},
		}
		target := -10000
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("边界最大值：m=100,n=100", func(t *testing.T) {
		// 创建100x100的矩阵
		matrix := make([][]int, 100)
		value := 1
		for i := 0; i < 100; i++ {
			matrix[i] = make([]int, 100)
			for j := 0; j < 100; j++ {
				matrix[i][j] = value
				value++
			}
		}
		target := 5000
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("负数值存在", func(t *testing.T) {
		matrix := [][]int{
			{-10, -5, -3, -1},
			{0, 2, 4, 6},
			{8, 10, 12, 14},
		}
		target := -5
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("负数值不存在", func(t *testing.T) {
		matrix := [][]int{
			{-10, -5, -3, -1},
			{0, 2, 4, 6},
			{8, 10, 12, 14},
		}
		target := -7
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("重复值存在", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 10, 16, 20},
			{23, 30, 34, 60},
		}
		target := 10
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("矩阵中间值存在", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
			{61, 62, 63, 64},
		}
		target := 30
		got := searchMatrix(matrix, target)
		expected := true

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("矩阵中间值不存在", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
			{61, 62, 63, 64},
		}
		target := 25
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("空矩阵（边界情况）", func(t *testing.T) {
		matrix := [][]int{}
		target := 5
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("空行矩阵（边界情况）", func(t *testing.T) {
		matrix := [][]int{{}}
		target := 5
		got := searchMatrix(matrix, target)
		expected := false

		if got != expected {
			t.Errorf("searchMatrix(matrix, %d) = %v, 期望 %v", target, got, expected)
		}
	})

	t.Run("验证矩阵属性：每行递增", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}

		// 验证矩阵属性
		for i := 0; i < len(matrix); i++ {
			for j := 1; j < len(matrix[i]); j++ {
				if matrix[i][j] < matrix[i][j-1] {
					t.Errorf("矩阵不满足每行递增属性：行[%d]列[%d]=%d < 列[%d]=%d", i, j, matrix[i][j], j-1, matrix[i][j-1])
				}
			}
		}

		// 验证每行第一个大于前一行最后一个
		for i := 1; i < len(matrix); i++ {
			if matrix[i][0] <= matrix[i-1][len(matrix[i-1])-1] {
				t.Errorf("矩阵不满足行间属性：行[%d]第一个=%d <= 行[%d]最后一个=%d", i, matrix[i][0], i-1, matrix[i-1][len(matrix[i-1])-1])
			}
		}
	})
}
