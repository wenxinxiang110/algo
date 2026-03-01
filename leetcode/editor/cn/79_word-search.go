package main

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
//输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
//"SEE"
//输出：true
//
//
// 示例 3：
//
//
//输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
//"ABCB"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
//
// Related Topics 深度优先搜索 数组 字符串 回溯 矩阵 👍 2138 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) (exists bool) {
	if len(word) == 0 || len(board) == 0 {
		return false
	}

	//row, column := len(board), len(board[0])

	// 初始化已经访问过的数组
	visit := make([][]bool, len(board))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(board[0]))
	}

	var backtrace func(i, j int, k int) bool
	backtrace = func(i, j int, k int) bool {

		//	终止条件
		if i < 0 || i >= len(board) ||
			j < 0 || j >= len(board[0]) ||
			word[k] != board[i][j] || visit[i][j] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		visit[i][j] = true
		defer func() { visit[i][j] = false }()

		return backtrace(i-1, j, k+1) ||
			backtrace(i+1, j, k+1) ||
			backtrace(i, j-1, k+1) ||
			backtrace(i, j+1, k+1)
	}

	for i, row := range board {
		for j := range row {
			exists = backtrace(i, j, 0)
			if exists {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
