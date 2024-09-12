package main

//给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。
//
// 螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。
//
//
//
// 示例 1：
//
//
//输入：array = [[1,2,3],[8,9,4],[7,6,5]]
//输出：[1,2,3,4,5,6,7,8,9]
//
//
// 示例 2：
//
//
//输入：array  = [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]
//输出：[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]
//
//
//
//
// 限制：
//
//
// 0 <= array.length <= 100
// 0 <= array[i].length <= 100
//
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/
//
//
//
// Related Topics 数组 矩阵 模拟 👍 603 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return nil
	}
	var width, length = len(array), len(array[0])
	var result = make([]int, width*length)
	var idx = 0
	var (
		left   = 0
		right  = length - 1
		top    = 0
		bottom = width - 1
	)
	for {
		// 左上角出发开始遍历
		for i := left; i <= right; i++ {
			result[idx] = array[top][i]
			idx++
		}
		top++
		if top > bottom {
			break
		}

		// 右上角出发开始遍历
		for i := top; i <= bottom; i++ {
			result[idx] = array[i][right]
			idx++
		}
		right--
		if right < left {
			break
		}
		// 右下角出发遍历
		for i := right; i >= left; i-- {
			result[idx] = array[bottom][i]
			idx++
		}
		bottom--
		if top > bottom {
			break
		}

		// 左下角出发遍历
		for i := bottom; i >= top; i-- {
			result[idx] = array[i][left]
			idx++
		}
		left++
		if right < left {
			break
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
