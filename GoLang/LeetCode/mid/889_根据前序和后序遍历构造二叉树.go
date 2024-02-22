package mid

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return build_889(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
}

func build_889(preorder, postorder []int, preStart, preEnd, postStart, postEnd int) *TreeNode {
	if preStart > preEnd || postStart > postEnd {
		return nil
	}
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	if preStart+1 <= preEnd {
		leftVal := preorder[preStart+1]

		i := postStart
		for ; i <= postEnd-1 && postorder[i] != leftVal; i++ {
		}
		gap := i - postStart

		root.Left = build_889(preorder, postorder, preStart+1, preStart+1+gap, postStart, i)
		root.Right = build_889(preorder, postorder, preStart+1+gap+1, preEnd, i+1, postEnd-1)
	}
	return root
}
