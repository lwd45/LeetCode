package mid

import "math"

var ans int

func maxAncestorDiff(root *TreeNode) int {
	ans = 0
	dfs_1026(root, root.Val, root.Val)
	return ans
}

func dfs_1026(root *TreeNode, big, small int) {
	if root == nil {
		return
	}

	big = max(big, root.Val)
	small = min(small, root.Val)
	ans = max(ans, int(math.Abs(float64(big-small))))
	dfs_1026(root.Right, big, small)
	dfs_1026(root.Left, big, small)
}
