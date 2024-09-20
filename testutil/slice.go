package testutil

import (
	"github.com/samber/lo"
)

// SliceEqual 判断数组相等,忽略顺序
func SliceEqual[T comparable](s1, s2 []T) bool {

	m1 := lo.CountValues(s1)
	m2 := lo.CountValues(s2)

	if len(m1) != len(m2) {
		return false
	}
	for v, count := range m1 {
		if m2[v] != count {
			return false
		}
	}
	return true
}
