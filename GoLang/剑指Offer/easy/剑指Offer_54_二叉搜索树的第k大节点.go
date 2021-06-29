package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var ans int
	helper(root, &k, &ans)
	return ans
}

func helper(root *TreeNode, k *int, ans *int) {
	if root == nil || *k < 1 {
		return
	}

	helper(root.Right, k, ans)
	if *k == 1 {
		*ans = root.Val
	}
	*k--
	helper(root.Left, k, ans)
}
