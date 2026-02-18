package main

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 矩阵 模拟 👍 2068 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	var left, right = 0, len(matrix[0]) - 1
	var top, bottom = 0, len(matrix) - 1
	for left <= right || top <= bottom {
		for i := left; i <= right && top <= bottom; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom && left <= right; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		for i := right; i >= left && top <= bottom; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top && left <= right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	//if left == right && top == bottom {
	//	res = append(res, matrix[top][left])
	//}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
