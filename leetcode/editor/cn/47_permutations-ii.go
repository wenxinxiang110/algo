package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 排序 👍 1771 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) (output [][]int) {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	visit := make([]bool, len(nums))

	var backtrace func(depth int, path []int)
	backtrace = func(depth int, path []int) {
		if depth == len(nums) {
			output = append(output,
				append([]int(nil), path...))
			return
		}

		for i, v := range nums {
			// 判断是否已经访问
			if visit[i] ||
				(i > 0 && v == nums[i-1] && !visit[i-1]) {
				continue
			}
			path = append(path, v)
			visit[i] = true

			fmt.Println("回朔之前,path:", path, "visit:", visit)

			backtrace(depth+1, path)

			visit[i] = false
			path = path[:len(path)-1]

			fmt.Println("回朔之后,path:", path, "visit:", visit)
		}

	}

	backtrace(0, nil)

	return
}

//leetcode submit region end(Prohibit modification and deletion)
