package mid

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	return f(&preorder, &inorder, 0, len(preorder)-1, 0, len(inorder)-1)
//}
//
//func f(preOrder, inOrder *[]int, pStart, pEnd, iStart, iEnd int) (root *TreeNode) {
//	if pStart > pEnd || iStart > iEnd {
//		return
//	}
//
//	root = &TreeNode{
//		Val:   (*preOrder)[pStart],
//		Left:  nil,
//		Right: nil,
//	}
//
//	piv := iStart
//	for i := iStart; i <= iEnd; i++ {
//		if (*inOrder)[i] == (*preOrder)[pStart] {
//			piv = i
//			break
//		}
//	}
//	lens := piv - iStart
//	root.Left = f(preOrder, inOrder, pStart+1, pStart+lens, iStart, piv-1)
//	root.Right = f(preOrder, inOrder, pStart+lens+1, pEnd, piv+1, iEnd)
//	return
//}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
//}
//
//func build(preorder, inorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
//	if pStart > pEnd || iStart > iEnd {
//		return nil
//	}
//
//	rootVal := preorder[pStart]
//	root := &TreeNode{Val: rootVal}
//
//	i := iStart
//	for ; i <= iEnd && inorder[i] != rootVal; i++ {
//	}
//	gap := i - iStart
//
//	root.Left = build(preorder, inorder, pStart+1, pStart+gap, iStart, i-1)
//	root.Right = build(preorder, inorder, pStart+gap+1, pEnd, i+1, iEnd)
//	return root
//}
