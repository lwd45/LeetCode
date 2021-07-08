package easy

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper28(root.Left, root.Right)
}

func helper28(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || (left.Val != right.Val) {
		return false
	}
	return helper28(left.Left, right.Right) && helper28(left.Right, right.Left)
}
