package main

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 3256 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) (output [][]int) {
	if len(nums) == 0 {
		return nil
	}
	// 1. 初始化 output数组和已经使用的元素数组
	visit := make([]bool, len(nums))

	var backtrace func(nums []int, depth int, path []int)
	backtrace = func(nums []int, depth int, path []int) {
		if depth == len(nums) {
			// clone path
			tmp := make([]int, len(path))
			copy(tmp, path)
			output = append(output, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visit[i] {
				continue
			}
			// 之前没有使用过的排列
			path = append(path, nums[i])
			visit[i] = true

			//fmt.Println("回朔之前,depth:", depth, "path:", path)

			backtrace(nums, depth+1, path)
			// 回退
			visit[i] = false
			path = path[:len(path)-1]
			//fmt.Println("回朔之后,depth:", depth, "path:", path)

		}
	}

	backtrace(nums, 0, nil)
	return output
}

//leetcode submit region end(Prohibit modification and deletion)
