package main

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
//
// Related Topics 位运算 数组 回溯 👍 2596 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) (output [][]int) {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var backtrace func(int, []int)
	backtrace = func(start int, current []int) {
		// 添加当前子集到结果
		tmp := make([]int, len(current))
		copy(tmp, current)
		output = append(output, tmp)

		// 继续添加元素
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrace(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrace(0, []int{})
	return
}

//leetcode submit region end(Prohibit modification and deletion)
