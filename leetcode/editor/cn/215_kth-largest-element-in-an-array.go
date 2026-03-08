package main

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2928 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	// 使用最小堆来维护前k大的元素： 在这个最小堆里，堆顶元素直接就是第k大的元素
	minHeap := make([]int, 0, k)

	var heapUp = func(index int) {
		for index > 0 {
			parent := (index - 1) / 2
			if minHeap[parent] <= minHeap[index] {
				break
			}
			minHeap[index], minHeap[parent] = minHeap[parent], minHeap[index]
			index = parent
		}
	}

	var heapDown = func(index int) {
		for {
			left := 2*index + 1
			right := 2*index + 2

			var smallest int = index
			// 找到两孩子的最小值，并比较
			if left < len(minHeap) && minHeap[left] < minHeap[index] {
				smallest = left
			}
			if right < len(minHeap) && minHeap[right] < minHeap[smallest] {
				smallest = right
			}
			// 已经比两个孩子小了，不需要再操作了
			if smallest == index {
				break
			}
			// swap
			minHeap[index], minHeap[smallest] = minHeap[smallest], minHeap[index]
			index = smallest
		}
	}

	for _, num := range nums {
		if len(minHeap) < k {
			minHeap = append(minHeap, num)
			heapUp(len(minHeap) - 1)
		} else if num > minHeap[0] {
			// 只有当新元素比堆顶元素大时，才替换堆顶元素
			minHeap[0] = num
			heapDown(0)
		}
	}

	return minHeap[0]
}

// 堆排序下降： 和子节点比较，如果比子节点小，交换
func heapifyDown(heap []int, i int, i2 int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < i2 && heap[left] > heap[largest] {
			largest = left
		}
		if right < i2 && heap[right] > heap[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		heap[i], heap[largest] = heap[largest], heap[i]
		i = largest
	}
}

// 堆排序上升： 和父节点比较，如果比父节点大，交换
func heapifyUp(heap []int, i int) {

	for i > 0 {
		parent := (i - 1) / 2
		if heap[i] <= heap[parent] {
			break
		}
		heap[i], heap[parent] = heap[parent], heap[i]
		i = parent
	}
}

//leetcode submit region end(Prohibit modification and deletion)
