package dp

// BasePackage 基础背包问题
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.md
func BasePackage(items []Item, maxWeight int) (maxValue int) {
	if len(items) == 0 || maxWeight == 0 {
		return 0
	}
	// 滚动数组实现，用一维数组存储压缩存储
	// 因为这个背包每次其实只需要读取上一行的数据，所以其实没必要存储整个完整的 dp表格
	dp := make([]int, maxWeight+1)
	for i := 0; i < len(items); i++ {
		// 如果是滚动数组方式的，需要用倒序遍历，避免一个物品重复放入
		for j := maxWeight; j >= items[i].Weight; j-- {
			dp[j] = max(dp[j], dp[j-items[i].Weight]+items[i].Value)
		}
	}
	//fmt.Println(dp)
	return dp[maxWeight]
}

type Item struct {
	Weight int
	Value  int
}
