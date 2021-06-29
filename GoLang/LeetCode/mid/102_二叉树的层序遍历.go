package mid

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	var level []*TreeNode
	level = append(level, root)

	for len(level) > 0 {
		size := len(level)
		var one []int
		for i := 0; i < size; i++ {
			poll := level[0]
			level = level[1:]

			one = append(one, poll.Val)
			if poll.Left != nil {
				level = append(level, poll.Left)
			}
			if poll.Right != nil {
				level = append(level, poll.Right)
			}
		}
		ret = append(ret, one)
	}
	return ret
}
