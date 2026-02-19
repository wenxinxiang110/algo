package main

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 二分查找 分治 矩阵 👍 1770 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	return matrixSearchZ(matrix, target)
}

// 使用暴力的方法
func searchMatrixSlow(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}

// 使用二分查找
func searchMatrixBinary(matrix [][]int, target int) bool {
	for _, rows := range matrix {

		var start, end = 0, len(rows) - 1
		for start <= end {
			var mid = start + (end-start)/2
			if rows[mid] == target {
				return true
			} else if rows[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return false
}

// 使用Z字形搜索,从右上角开始查找，如果 > target, 向左移动; < target,向下移动
func matrixSearchZ(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var row, col = 0, len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
