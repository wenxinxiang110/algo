package tree

import (
	"container/list"
)

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return IsMirrorRecursive(root.Left, root.Right)

	//return IsMirrorIterate(root.Left, root.Right)
}

/**
 * 递归检查两颗树是否镜像对称
 *
 * param: *TreeNode root1 第一颗树根节点
 * param: *TreeNode root2 第二颗树根节点
 * return: bool 是否镜像对称
 */
func IsMirrorRecursive(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil ||
		root1.Val != root2.Val {
		return false
	}

	return IsMirrorRecursive(root1.Left, root2.Right) &&
		IsMirrorRecursive(root1.Right, root2.Left)
}

/**
 * 递归检查两颗树是否镜像对称
 * todo:醉了，连个正常点的队列都没有，不会整
 *
 * param: *TreeNode root1 第一颗树根节点
 * param: *TreeNode root2 第二颗树根节点
 * return: bool	是否镜像对称
 */
func IsMirrorIterate(root1 *TreeNode, root2 *TreeNode) bool {
	queue := list.New()

	queue.PushBack(root1)
	queue.PushBack(root2)

	for queue.Len() != 0 {
		// 从队列中pull最前边的元素
		e1 := queue.Front()
		queue.Remove(e1)

		e2 := queue.Front()
		queue.Remove(e2)

		if e1.Value == nil && e2.Value == nil {
			continue
		}

		if e1.Value == nil || e2.Value == nil {
			return false
		}

		// 类型强转，吐了
		node1, ok := e1.Value.(*TreeNode)
		node2, ok2 := e2.Value.(*TreeNode)
		if !ok && !ok2 {
			continue
		}
		if !ok || !ok2 {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		queue.PushBack(node1.Left)
		queue.PushBack(node2.Right)
		queue.PushBack(node1.Right)
		queue.PushBack(node2.Left)

	}

	return true
}
