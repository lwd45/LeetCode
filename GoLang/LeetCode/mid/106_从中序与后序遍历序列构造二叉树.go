package mid

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build_106(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build_106(inorder, postorder []int, iStart, iEnd, pStart, pEnd int) *TreeNode {
	if iStart > iEnd || pStart > pEnd {
		return nil
	}

	rootVal := postorder[pEnd]
	root := &TreeNode{Val: rootVal}

	i := iStart
	for ; i <= iEnd && inorder[i] != rootVal; i++ {
	}
	gap := i - iStart

	root.Left = build_106(inorder, postorder, iStart, i-1, pStart, pStart+gap-1)
	root.Right = build_106(inorder, postorder, i+1, iEnd, pStart+gap, pEnd-1)

	return root
}
