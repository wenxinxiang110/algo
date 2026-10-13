package main

import (
	"sort"
)

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//
//
// 示例 1：
//
//
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。
//
// 示例 2：
//
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3：
//
//
//输入: candidates = [2], target = 1
//输出: []
//
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// candidates 的所有元素 互不相同
// 1 <= target <= 40
//
//
// Related Topics 数组 回溯 👍 3166 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) (output [][]int) {
	if len(candidates) == 0 {
		return nil
	}
	// 按照升序排列
	sort.Ints(candidates)

	var backtrace func(start int, path []int, target int)

	backtrace = func(start int, path []int, target int) {
		// 目标值小于0，说明当前路径不可行
		if target < 0 {
			return
		}
		// 找到合适的组合
		if target == 0 {
			output = append(output, append([]int{}, path...))
			return
		}

		// 从start开始搜索，允许重复使用同一个数字
		for i := start; i < len(candidates); i++ {
			// 如果当前数字已经大于目标值，由于数组已排序，后续数字都会更大，可以直接返回
			if candidates[i] > target {
				break
			}

			// 选择当前数字
			path = append(path, candidates[i])

			// 继续从i开始搜索（允许重复使用当前数字），目标值减去当前数字
			backtrace(i, path, target-candidates[i])

			// 回溯，撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrace(0, nil, target)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
