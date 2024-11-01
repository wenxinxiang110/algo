package dp

import (
	"fmt"
)

// BasePackage 基础背包问题
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.md
func BasePackage(items []Item, maxWeight int) (maxValue int) {
	if len(items) == 0 || maxWeight == 0 {
		return 0
	}
	dp := make([][]int, len(items))
	// 初始化
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}
	for i := items[0].Weight; i < maxWeight; i++ {
		dp[0][i] = items[0].Value
	}
	for i := 1; i < len(items); i++ {
		for j := 1; j <= maxWeight; j++ {
			if j < items[i].Weight {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i].Weight]+items[i].Value)
			}
		}
	}
	fmt.Println(dp)
	return dp[len(items)-1][maxWeight]
}

type Item struct {
	Weight int
	Value  int
}
