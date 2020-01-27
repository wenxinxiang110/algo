package DP

import (
	"fmt"
)

func MaxSubArray(nums []int) int {
	n := len(nums)
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)

		m[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			m[i][j] = m[i][j-1] + nums[j]
		}
	}
	max := m[0][0]
	for i, v := range m {
		for j := i; j < n; j++ {
			if m[i][j] > max {
				max = m[i][j]
			}
		}
		fmt.Println(v)
	}
	return max
}
