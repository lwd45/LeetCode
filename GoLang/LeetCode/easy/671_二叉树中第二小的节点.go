package easy

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	ans := -1
	dfs_671(root, root.Val, &ans)
	return ans
}

func dfs_671(root *TreeNode, rootVal int, ans *int) {
	if root == nil {
		return
	}
	if root.Val > *ans && *ans != -1 {
		return
	}
	if root.Val > rootVal && (*ans == -1 || root.Val < *ans) {
		*ans = root.Val
	}
	dfs_671(root.Left, rootVal, ans)
	dfs_671(root.Right, rootVal, ans)
}
