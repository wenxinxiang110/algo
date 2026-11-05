package heap

type heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func (h *heap[T]) Push(val T) {
	h.data = append(h.data, val)
	h.heapUp(len(h.data) - 1)
}

func (h *heap[T]) Pop() (val T) {
	if len(h.data) == 0 {
		return
	}
	val = h.data[0]

	// 将末尾元素移到堆顶，并缩短切片，再对堆顶下沉
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.heapDown(0)
	}
	return
}

func (h *heap[T]) Top() (val T) {
	if len(h.data) == 0 {
		return
	}
	return h.data[0]
}

// Len 返回元素数量
func (h *heap[T]) Len() int {
	return len(h.data)
}

// 核心操作，把 idx 位置的元素上浮，基于 less比较
func (h *heap[T]) heapUp(cur int) {
	for cur > 0 {
		parent := (cur - 1) / 2
		if !h.less(h.data[cur], h.data[parent]) {
			break
		}
		// 当前元素比较小
		h.data[cur], h.data[parent] = h.data[parent], h.data[cur]
		cur = parent
	}
}

// 核心操作2:下沉，和自己比较小的孩子比较
func (h *heap[T]) heapDown(cur int) {
	for cur < len(h.data)-1 {
		left := 2*cur + 1
		right := 2*cur + 2

		smallest := cur
		if left < len(h.data) && h.less(h.data[left], h.data[smallest]) {
			smallest = left
		}
		if right < len(h.data) && h.less(h.data[right], h.data[smallest]) {
			smallest = right
		}
		if smallest == cur {
			break
		}
		// 交换
		h.data[cur], h.data[smallest] = h.data[smallest], h.data[cur]

		cur = smallest
	}
}
