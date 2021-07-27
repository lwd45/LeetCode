package mid

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[int]*TreeNode)
	findParent(root, parents)
	var ans []int
	findAns(target, nil, &ans, parents, 0, k)
	return ans
}

func findAns(target, from *TreeNode, ans *[]int, parents map[int]*TreeNode, deep, k int) {
	if target == nil {
		return
	}
	if deep == k {
		*ans = append(*ans, target.Val)
		return
	}
	if target.Left != from {
		findAns(target.Left, target, ans, parents, deep+1, k)
	}
	if target.Right != from {
		findAns(target.Right, target, ans, parents, deep+1, k)
	}
	if parents[target.Val] != from {
		findAns(parents[target.Val], target, ans, parents, deep+1, k)
	}
}

func findParent(root *TreeNode, parents map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		parents[root.Left.Val] = root
		findParent(root.Left, parents)
	}
	if root.Right != nil {
		parents[root.Right.Val] = root
		findParent(root.Right, parents)
	}
}
