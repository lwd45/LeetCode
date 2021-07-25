package mid

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	var layer []*TreeNode
	layer = append(layer, root)
	for len(layer) > 0 {
		oneAns := []int{}
		for i := len(layer); i > 0; i-- {
			pop := layer[0]
			layer = layer[1:]
			oneAns = append(oneAns, pop.Val)
			if pop.Left != nil {
				layer = append(layer, pop.Left)
			}
			if pop.Right != nil {
				layer = append(layer, pop.Right)
			}
		}
		ans = append(ans, oneAns)
	}
	return ans
}
