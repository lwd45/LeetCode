package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ans := 0
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, ans)
	right := helper(root.Right, ans)
	if left > right {
		*ans += left - right
	} else {
		*ans += right - left
	}
	return left + right + root.Val
}
