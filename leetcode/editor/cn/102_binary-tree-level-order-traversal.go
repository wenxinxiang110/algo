package main

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 2249 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// pop queue all element
		elements := queue[:]
		queue = nil
		var levelResult = make([]int, 0, len(elements))
		for _, element := range elements {
			//set into queue
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}

			// record
			levelResult = append(levelResult, element.Val)
		}
		res = append(res, levelResult)
	}
	return

}

//leetcode submit region end(Prohibit modification and deletion)
