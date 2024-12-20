package main

import (
	"container/list"

	"github.com/NothingXiang/algo/leetcode/monotonic_queue"
)

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2907 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	//先将前k的元素放入队列
	var queue = list.New()
	for i := 0; i < k; i++ {
		// 将队首元素移除 替换queue.Push(nums[i])
		for queue.Back() != nil && queue.Back().Value.(int) < nums[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(nums[i])
	}
	res := make([]int, len(nums)-k+1)
	res[0] = queue.Front().Value.(int)

	for i := k; i < len(nums); i++ {
		// 将队首元素移除 替换queue.Pop(nums[i-k])
		if queue.Front().Value.(int) == nums[i-k] {
			queue.Remove(queue.Front())
		}

		// 添加到队尾 替换 queue.Push(nums[i])
		for queue.Back() != nil && queue.Back().Value.(int) < nums[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(nums[i])

		res[i-k+1] = queue.Front().Value.(int)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func maxV1(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	//先将前k的元素放入队列
	var queue = monotonic_queue.New[int]()
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res := make([]int, len(nums)-k+1)
	res[0] = queue.Max()
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res[i-k+1] = queue.Max()
	}
	return res
}
