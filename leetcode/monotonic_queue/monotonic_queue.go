package monotonic_queue

import (
	"cmp"
	"container/list"
)

// Queue 单调队列
type Queue[T cmp.Ordered] interface {
	// Push 添加元素 val
	Push(val T)
	// Pop 如果队尾元素是 val 则删除
	Pop(val T)
	// Max 返回队列中最大的元素
	Max() T
}

func New[T cmp.Ordered]() Queue[T] {
	return &queue[T]{List: list.New()}
}

type queue[T cmp.Ordered] struct {
	List *list.List
}

func (q *queue[T]) Push(val T) {
	for q.List.Back() != nil && q.List.Back().Value.(T) < val {
		q.List.Remove(q.List.Back())
	}
	q.List.PushBack(val)
}

func (q *queue[T]) Pop(val T) {
	if q.List.Back().Value == val {
		q.List.Remove(q.List.Back())
	}
}

// Max 返回队列中最大的元素
// 单调队列中，队头就是最大的元素
func (q *queue[T]) Max() (val T) {
	front := q.List.Front()
	if front != nil {
		return front.Value.(T)
	}
	return
}
