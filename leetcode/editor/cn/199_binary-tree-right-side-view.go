package main

//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,2,3,null,5,null,4]
//
//
// 输出：[1,3,4]
//
// 解释：
//
//
//
// 示例 2：
//
//
// 输入：root = [1,2,3,4,null,null,null,5]
//
//
// 输出：[1,3,4,5]
//
// 解释：
//
//
//
// 示例 3：
//
//
// 输入：root = [1,null,3]
//
//
// 输出：[1,3]
//
// 示例 4：
//
//
// 输入：root = []
//
//
// 输出：[]
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,100]
//
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1293 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) (view []int) {
	// 层次序遍历
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// pop queue all element
		elems := queue[:]
		queue = nil

		view = append(view, elems[len(elems)-1].Val)

		// set into queue
		for _, elem := range elems {
			if elem.Left != nil {
				queue = append(queue, elem.Left)
			}
			if elem.Right != nil {
				queue = append(queue, elem.Right)
			}
		}

	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
