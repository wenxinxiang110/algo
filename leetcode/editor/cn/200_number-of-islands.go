package main

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ['1','1','1','1','0'],
//  ['1','1','0','1','0'],
//  ['1','1','0','0','0'],
//  ['0','0','0','0','0']
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ['1','1','0','0','0'],
//  ['1','1','0','0','0'],
//  ['0','0','1','0','0'],
//  ['0','0','0','1','1']
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2931 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	return numIslandsDFS(grid)

}

const island = '1'

func numIslandsDFS(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == island {
				count++
				dfsGrid(grid, i, j)
			}
		}
	}
	return
}

// dfs标记 i,j位置的岛屿
func dfsGrid(grid [][]byte, row int, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) || grid[row][column] != island {
		return
	}
	// 已经遍历过的，标记为0
	grid[row][column] = '0'

	dfsGrid(grid, row-1, column)
	dfsGrid(grid, row+1, column)
	dfsGrid(grid, row, column-1)
	dfsGrid(grid, row, column+1)

}

//leetcode submit region end(Prohibit modification and deletion)
