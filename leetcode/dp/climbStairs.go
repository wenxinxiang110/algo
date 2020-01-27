package DP

// 爬楼梯
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	mem := make([]int, n)

	mem[0] = 1
	mem[1] = 2
	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n-1]
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		if i == 0 {
			m[i][0] = cost[0]
		} else {
			m[i][0] = m[i-1][0] + cost[i]
		}
		m[i][i] = cost[i]
		for j := 1; j < i; j++ {
			m[i][j] = m[i][j-1]
		}
	}
	max := 0
	for _, v := range m {
		for _, vv := range v {
			if vv > max {
				max = vv
			}
		}
	}
	return max
}
func sum(i, j int, cost []int) int {
	r := 0
	for ; i <= j; i++ {
		r += cost[i]
	}
	return r
}

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
