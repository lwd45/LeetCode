package mid

func ZigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0, 10)
	if root == nil {
		return ans
	}
	nodeList := make([]*TreeNode, 0, 20)
	nodeList = append(nodeList, root)

	l2r := true
	for len(nodeList) > 0 {
		one := make([]int, 0, 10)
		for size := len(nodeList); size > 0; size-- {
			poll := nodeList[0]
			nodeList = nodeList[1:]
			if poll.Left != nil {
				nodeList = append(nodeList, poll.Left)
			}
			if poll.Right != nil {
				nodeList = append(nodeList, poll.Right)
			}

			if l2r {
				one = append(one, poll.Val)
			} else {
				if len(one) > 0 {
					one = append(one, 0)
					copy(one[1:], one[:])
					one[0] = poll.Val
				} else {
					one = append(one, poll.Val)
				}
			}
		}
		l2r = !l2r
		ans = append(ans, one)
	}
	return ans
}
