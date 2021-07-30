package mid

func pathSum(root *TreeNode, target int) [][]int {
	ans, one := make([][]int, 0), make([]int, 0)
	sum := 0

	dfs_34(root, &ans, &one, &sum, target)
	return ans
}

func dfs_34(root *TreeNode, ans *[][]int, one *[]int, sum *int, target int) {
	if root == nil {
		return
	}
	*sum += root.Val
	*one = append(*one, root.Val)
	if root.Left == nil && root.Right == nil && *sum == target {
		oneAns := make([]int, len(*one))
		copy(oneAns, *one)
		*ans = append(*ans, oneAns)
	}
	dfs_34(root.Left, ans, one, sum, target)
	dfs_34(root.Right, ans, one, sum, target)
	*sum -= root.Val
	*one = (*one)[:len(*one)-1]
}
