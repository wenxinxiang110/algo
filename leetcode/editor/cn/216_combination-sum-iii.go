package main

//找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
//
// 只使用数字1到9
// 每个数字 最多使用一次
//
//
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
//
//
// 示例 1:
//
//
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//解释:
//1 + 2 + 4 = 7
//没有其他符合的组合了。
//
// 示例 2:
//
//
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
//解释:
//1 + 2 + 6 = 9
//1 + 3 + 5 = 9
//2 + 3 + 4 = 9
//没有其他符合的组合了。
//
// 示例 3:
//
//
//输入: k = 4, n = 1
//输出: []
//解释: 不存在有效的组合。
//在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
//
//
//
// 提示:
//
//
// 2 <= k <= 9
// 1 <= n <= 60
//
//
// Related Topics 数组 回溯 👍 997 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) (comb [][]int) {

	var backtrace func(path []int, start, k, target int)
	// path: 保存前面已经搜索过的结果
	// start: 这次循环可以从哪里开始填充值
	// k: 还需要填充值的个数
	// n: 还需要填充值的大小
	backtrace = func(path []int, start, k, target int) {
		// 终止条件: 只剩下一个数，要么有要么没有，一定要结束了
		if k == 1 {
			// 终止条件: 最后一个数大于 start
			if target > 9 {
				return
			}
			if target >= start {
				path = append(path, target)
				var res = make([]int, len(path))
				copy(res, path)
				comb = append(comb, res)
			}
			return
		}
		for i := start; i <= target && i <= 9; i++ {
			path = append(path, i)
			backtrace(path, i+1, k-1, target-i)
			path = path[:len(path)-1]
		}
	}

	backtrace(make([]int, 0, k), 1, k, n)
	return

}

//leetcode submit region end(Prohibit modification and deletion)
