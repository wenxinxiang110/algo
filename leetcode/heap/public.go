package heap

import (
	"cmp"
)

// Interface 是堆的通用行为接口，任何实现了 Push/Pop/Top/Len 的类型都可被视为堆。
// 借助该接口，调用方可以以“堆”的多态方式使用不同实现（最小堆、最大堆等），
// 而无需关心底层具体类型。
type Interface[T any] interface {
	Push(val T)
	Pop() T
	Top() T
	Len() int
}

func MinHeap[T cmp.Ordered]() Interface[T] {
	return &heap[T]{
		less: MinCompare[T],
	}
}

func MaxHeap[T cmp.Ordered]() Interface[T] {
	return &heap[T]{
		less: MaxCompare[T],
	}
}

func New[T any](less func(a, b T) bool) Interface[T] {
	return &heap[T]{less: less}
}

// MinCompare min比较
func MinCompare[T cmp.Ordered](a, b T) bool {
	return a < b
}

// MaxCompare min比较
func MaxCompare[T cmp.Ordered](a, b T) bool {
	return a > b
}
