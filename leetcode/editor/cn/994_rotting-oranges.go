package main

//在给定的 m x n 网格
// grid 中，每个单元格可以有以下三个值之一：
//
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
//
//
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4
//
//
// 示例 2：
//
//
//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
//
//
// 示例 3：
//
//
//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] 仅为 0、1 或 2
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 1127 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	type Point struct {
		Row, Column int
	}

	queue := []Point{}

	var freshCount = 0
	// 第一次遍历： 烂橘子添加到 queue 中，记录好的橘子数量
	for i, row := range grid {
		for j, orange := range row {
			if orange == fresh {
				freshCount++
			} else if orange == rotting {
				queue = append(queue, Point{i, j})
			}
		}
	}
	if freshCount == 0 {
		return 0
	}

	loop := -1
	for len(queue) != 0 {
		// pop all
		points := queue[:]
		queue = []Point{}

		for _, point := range points {
			//上下左右
			if point.Row > 0 && grid[point.Row-1][point.Column] == fresh {
				grid[point.Row-1][point.Column] = rotting
				freshCount--
				queue = append(queue, Point{Row: point.Row - 1, Column: point.Column})
			}
			if point.Row < len(grid)-1 && grid[point.Row+1][point.Column] == fresh {
				grid[point.Row+1][point.Column] = rotting
				freshCount--
				queue = append(queue, Point{Row: point.Row + 1, Column: point.Column})
			}

			if point.Column > 0 && grid[point.Row][point.Column-1] == fresh {
				grid[point.Row][point.Column-1] = rotting
				freshCount--
				queue = append(queue, Point{point.Row, point.Column - 1})
			}
			if point.Column < len(grid[0])-1 && grid[point.Row][point.Column+1] == fresh {
				grid[point.Row][point.Column+1] = rotting
				freshCount--
				queue = append(queue, Point{point.Row, point.Column + 1})
			}
		}

		loop++
	}
	if freshCount > 0 {
		return -1
	}
	return loop
}

const (
	fresh   = 1
	rotting = 2
)

//leetcode submit region end(Prohibit modification and deletion)
