package mid

func levelOrderIII(root *TreeNode) [][]int {
	ans, layer, reverse := make([][]int, 0), make([]*TreeNode, 0), false
	if root == nil {
		return ans
	}
	layer = append(layer, root)

	for len(layer) > 0 {
		oneAns := make([]int, 0)
		for i := len(layer); i > 0; i-- {
			pop := layer[0]
			oneAns = append(oneAns, pop.Val)
			layer = layer[1:]
			if pop.Right != nil {
				layer = append(layer, pop.Right)
			}
			if pop.Left != nil {
				layer = append(layer, pop.Left)
			}
		}
		if reverse {
			for i, j := 0, len(oneAns)-1; i < j; i, j = i+1, j-1 {
				temp := oneAns[i]
				oneAns[i] = oneAns[j]
				oneAns[j] = temp
			}
		}
		reverse = !reverse
		ans = append(ans, oneAns)
	}
	return ans
}
