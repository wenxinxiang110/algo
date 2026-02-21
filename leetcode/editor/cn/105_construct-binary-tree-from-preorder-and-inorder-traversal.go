package main

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。
//
//
//
// 示例 1:
//
//
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//
//
//
//
// 提示:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2680 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历的第一个肯定是根节点
	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]

	// 在中序遍历中找到根节点
	var leftTreeCount = 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			leftTreeCount = i
			break
		}
	}
	if leftTreeCount > 0 {
		root.Left = buildTree(preorder[:leftTreeCount], inorder[:leftTreeCount])
	}
	if leftTreeCount < len(inorder)-1 {
		root.Right = buildTree(preorder[leftTreeCount:], inorder[leftTreeCount+1:])
	}
	return root

}

//leetcode submit region end(Prohibit modification and deletion)
