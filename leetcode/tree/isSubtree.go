package tree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == t {
		return true
	}

	if t == nil {
		return true
	}

	if s == nil {
		return false
	}

	if s.Val == t.Val {
		if isSameTree(s, t) {
			return true
		}
	}
	return isSameTree(s.Left, t) || isSameTree(s.Right, t)
}
