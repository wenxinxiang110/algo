package treenode

// 100. 相同的树
// 给定两个二叉树，编写一个函数来检验它们是否相同。
// 穷举
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}

	return false
}
