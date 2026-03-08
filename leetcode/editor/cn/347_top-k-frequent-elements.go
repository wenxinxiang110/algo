package main

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,1,2,2,3], k = 2
//
//
// 输出：[1,2]
//
// 示例 2：
//
//
// 输入：nums = [1], k = 1
//
//
// 输出：[1]
//
// 示例 3：
//
//
// 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2
//
//
// 输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 2143 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	// 统计每个元素的频率
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}

	// 创建最小堆，堆中存储的是元素值，但比较基于频率
	heap := make([]int, 0, k)

	// 堆上浮操作：基于频率比较
	var heapUp = func(idx int) {
		for idx > 0 {
			parent := (idx - 1) / 2
			// 比较频率：堆顶应该是最小频率的元素
			if frequencies[heap[parent]] <= frequencies[heap[idx]] {
				break
			}
			heap[parent], heap[idx] = heap[idx], heap[parent]
			idx = parent
		}
	}

	// 堆下沉操作：基于频率比较
	var heapDown = func(idx int) {
		for {
			left := 2*idx + 1
			right := 2*idx + 2
			smallest := idx

			if left < len(heap) && frequencies[heap[left]] < frequencies[heap[smallest]] {
				smallest = left
			}
			if right < len(heap) && frequencies[heap[right]] < frequencies[heap[smallest]] {
				smallest = right
			}

			if smallest == idx {
				break
			}
			heap[idx], heap[smallest] = heap[smallest], heap[idx]
			idx = smallest
		}
	}

	// 遍历频率表中的所有元素
	for v := range frequencies {
		if len(heap) < k {
			// 堆未满，直接添加并上浮
			heap = append(heap, v)
			heapUp(len(heap) - 1)
		} else if frequencies[v] > frequencies[heap[0]] {
			// 新元素的频率大于堆顶频率，替换堆顶并下沉
			heap[0] = v
			heapDown(0)
		}
		// 如果频率小于等于堆顶频率，忽略该元素
	}

	return heap
}

//leetcode submit region end(Prohibit modification and deletion)
