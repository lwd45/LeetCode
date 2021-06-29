package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	helper27(root)
	return root
}

func helper27(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil || root.Right != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		helper27(root.Left)
		helper27(root.Right)
	}
}
