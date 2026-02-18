package main

import (
	"sort"
)

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 示例 3：
//
//
//输入：intervals = [[4,7],[1,4]]
//输出：[[1,7]]
//解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
//
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics 数组 排序 👍 2719 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) (merged [][]int) {
	if len(intervals) == 0 {
		return
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		top := merged[len(merged)-1]
		if intervals[i][0] <= top[1] {
			// try merge
			top[1] = max(top[1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)
