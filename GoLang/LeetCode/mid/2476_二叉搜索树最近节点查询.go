package mid

import "sort"

func closestNodes(root *TreeNode, queries []int) [][]int {
	var values []int

	var dfs_2476 func(node *TreeNode)
	dfs_2476 = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs_2476(root.Left)
		values = append(values, root.Val)
		dfs_2476(root.Right)
	}
	dfs_2476(root)

	res := make([][]int, len(queries))
	for i, query := range queries {
		max, min := -1, -1
		idx := sort.SearchInts(values, query)

		if idx == len(values) { // 不存在
			min = values[idx-1]
			res[i] = []int{min, max}
			continue
		}

		if values[idx] == query {
			min = values[idx]
		} else if idx == 0 {
			min = -1
		} else {
			min = values[idx-1]
		}
		res[i] = []int{min, values[idx]}

	}
	return res
}
