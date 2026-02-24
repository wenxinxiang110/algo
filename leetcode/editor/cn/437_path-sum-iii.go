package main

//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
//
//
// 示例 2：
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,1000]
//
// -10⁹ <= Node.val <= 10⁹
// -1000 <= targetSum <= 1000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 2226 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	// 使用前缀和解法：更高效，时间复杂度O(n)
	return pathSumPrefix(root, targetSum)

	// 原来的递归解法：时间复杂度O(n^2)
	// return pathSumRecursive(root, targetSum)
}

// 前缀和解法：时间复杂度O(n)，空间复杂度O(n)
// 核心思想：记录从根节点到当前节点的路径和（前缀和），
// 然后检查是否存在前缀和等于当前前缀和减去targetSum的情况
func pathSumPrefix(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 记录前缀和，初始化为1
	prefixSumCount := map[int]int{
		0: 1,
	}

	return prefixSumDfs(root, 0, targetSum, prefixSumCount)
}

// 用 dfs获取前缀和
func prefixSumDfs(root *TreeNode, curSum, targetSum int, prefixSumCount map[int]int) (count int) {
	// 如果当前节点为空，处理结束
	if root == nil {
		return 0
	}

	curSum += root.Val
	count += prefixSumCount[curSum-targetSum]

	// 记录当前 sum
	prefixSumCount[curSum]++

	count += prefixSumDfs(root.Left, curSum, targetSum, prefixSumCount)
	count += prefixSumDfs(root.Right, curSum, targetSum, prefixSumCount)

	// 处理结束，把当前的前缀和去除
	prefixSumCount[curSum]--
	return
}

// 递归，遍历每一个节点， 时间复杂度和空间复杂度都是O(n^2)
func pathSumRecursive(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := rootSumRecursive(root, targetSum)
	count += pathSumRecursive(root.Left, targetSum)
	count += pathSumRecursive(root.Right, targetSum)
	return count
}

// 递归实现: 遍历每一个节点到子节点的实现： 时间复杂度比较差
func rootSumRecursive(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val == targetSum {
		count++
	}

	count += rootSumRecursive(root.Left, targetSum-root.Val)
	count += rootSumRecursive(root.Right, targetSum-root.Val)
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
