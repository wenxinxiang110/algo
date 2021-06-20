package tree

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftVal := root.Val + maxPathSum(root.Left)
	rightVal := root.Val + maxPathSum(root.Right)

	return Max(leftVal, rightVal)
}
