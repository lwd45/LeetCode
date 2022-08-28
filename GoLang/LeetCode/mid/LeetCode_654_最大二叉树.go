package mid

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTree2(nums, 0, len(nums)-1)
}

func constructMaximumBinaryTree2(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	maxIdx := start
	for i := start + 1; i <= end; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTree2(nums, start, maxIdx-1)
	root.Right = constructMaximumBinaryTree2(nums, maxIdx+1, end)
	return root
}
