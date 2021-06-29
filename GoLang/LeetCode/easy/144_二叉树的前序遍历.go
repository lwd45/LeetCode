package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	stack = append(stack, root)
	var ans []int
	if root == nil {
		return ans
	}
	for len(stack) > 0 {
		poll := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, poll.Val)

		if poll.Right != nil {
			stack = append(stack, poll.Right)
		}
		if poll.Left != nil {
			stack = append(stack, poll.Left)
		}
	}
	return ans
}
