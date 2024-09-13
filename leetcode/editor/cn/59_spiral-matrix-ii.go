package main

//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// Related Topics 数组 矩阵 模拟 👍 1330 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	var (
		left   = 0
		right  = n - 1
		top    = 0
		bottom = n - 1
	)
	var idx = 1
	for {
		// 左上角出发开始遍历
		for i := left; i <= right; i++ {
			matrix[top][i] = idx
			idx++
		}
		top++
		if top > bottom {
			break
		}

		// 右上角出发开始遍历
		for i := top; i <= bottom; i++ {
			matrix[i][right] = idx
			idx++
		}
		right--
		if right < left {
			break
		}
		// 右下角出发遍历
		for i := right; i >= left; i-- {
			matrix[bottom][i] = idx
			idx++
		}
		bottom--
		if top > bottom {
			break
		}

		// 左下角出发遍历
		for i := bottom; i >= top; i-- {
			matrix[i][left] = idx
			idx++
		}
		left++
		if right < left {
			break
		}
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
