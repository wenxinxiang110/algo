package main

import (
	"fmt"
	"strings"
)

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 2436 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) (ans [][]string) {
	// 皇后位置, <i,queens[i]>
	queens := make([]int, n)
	// 记录已经使用过的
	col := make([]bool, n)

	// 记录斜线是否可访问
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)

	var dfs func(int)
	dfs = func(row int) {
		// 终止条件，每一个皇后已经放置
		if row == n {
			ans = append(ans, outputQueens(queens))
			return
		}
		// 遍历每一列
		for column, ok := range col {
			// 没被访问过
			// row-column可能为负数溢出
			rc := row - column + (n - 1)

			if !ok && !diag1[row+column] && !diag2[rc] {
				queens[row] = column // 记录皇后位置，直接覆盖
				col[column] = true
				diag1[row+column] = true
				diag2[rc] = true

				fmt.Println("递归前，queens:", queens, "col:", col, "diag1:", diag1, "diag2:", diag2)

				dfs(row + 1)
				diag2[rc] = false
				diag1[row+column] = false
				col[column] = false

				fmt.Println("递归后，queens:", queens, "col:", col, "diag1:", diag1, "diag2:", diag2)
			}
		}
	}

	dfs(0)
	return

}

//关键点: 判断皇后是否可以放置：
// 1. 同一行：通过递归层数判断
// 2. 同一列：通过列数组记录
// 3. 同一对角线：对于 ↗ 方向的格子，行号加列号是不变的。对于 ↖ 方向的格子，行号减列号是不变的
//

// 辅助函数: 根据 queens 数组输出结果
// queens: 第 i 行的皇后对应的位置
func outputQueens(queens []int) (table []string) {
	const empty = "."
	const Queen = "Q"

	table = make([]string, len(queens))
	for rowIdx, colIdx := range queens {
		table[rowIdx] = strings.Repeat(empty, colIdx) + Queen +
			strings.Repeat(empty, len(queens)-colIdx-1)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
