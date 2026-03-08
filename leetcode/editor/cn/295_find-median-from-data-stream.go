package main

import (
	"container/heap"
)

//中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
//
// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
//
//
// 实现 MedianFinder 类:
//
//
// MedianFinder() 初始化 MedianFinder 对象。
// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10⁻⁵ 以内的答案将被接受。
//
//
// 示例 1：
//
//
//输入
//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//输出
//[null, null, null, 1.5, null, 2.0]
//
//解释
//MedianFinder medianFinder = new MedianFinder();
//medianFinder.addNum(1);    // arr = [1]
//medianFinder.addNum(2);    // arr = [1, 2]
//medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
//medianFinder.addNum(3);    // arr[1, 2, 3]
//medianFinder.findMedian(); // return 2.0
//
// 提示:
//
//
// -10⁵ <= num <= 10⁵
// 在调用 findMedian 之前，数据结构中至少有一个元素
// 最多 5 * 10⁴ 次调用 addNum 和 findMedian
//
//
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 1218 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type MedianFinder struct {
	//redblacktree.
	minHeap *intHeap
	maxHeap *intHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &intHeap{isMin: true},  // 所有大于等于 median 的数
		maxHeap: &intHeap{isMin: false}, // 所有小于 median 的数
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 首先插入到小顶堆
	//heap.Push(this.minHeap, num)
	//
	//expectedLen := (this.minHeap.Len() + this.maxHeap.Len() + 1) / 2
	//
	//adjusted := this.minHeap.Len() - expectedLen
	//for adjusted > 0 {
	//	val := heap.Pop(this.minHeap)
	//	heap.Push(this.maxHeap, val)
	//	adjusted--
	//}

	// 比较 medium的值，插入对应的堆
	if this.minHeap.Len() == 0 {
		heap.Push(this.minHeap, num)
		return
	}

	if num < this.minHeap.data[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	// 平衡两个堆
	if this.minHeap.Len() > this.maxHeap.Len()+1 {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
		return
	}

	if this.maxHeap.Len() > this.minHeap.Len() {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.data[0])
	}
	return float64(this.minHeap.data[0]+this.maxHeap.data[0]) / 2
}

// 辅助数据结构： heap
type intHeap struct {
	data []int
	//heap.Interface
	// 最小堆 or 最大堆
	isMin bool
}

func (hp *intHeap) Len() int {
	return len(hp.data)
}

func (hp *intHeap) Less(idx1, idx2 int) bool {
	if hp.isMin {
		return hp.data[idx1] < hp.data[idx2]
	}
	return hp.data[idx1] > hp.data[idx2]
}

func (hp *intHeap) Swap(i, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

func (hp *intHeap) Push(x any) {
	hp.data = append(hp.data, x.(int))
}

// Pop 弹出最后一个元素
// remove and return element Len() - 1.
func (hp *intHeap) Pop() any {
	idx := hp.Len() - 1
	val := hp.data[idx]
	hp.data = hp.data[:idx]
	return val
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)
