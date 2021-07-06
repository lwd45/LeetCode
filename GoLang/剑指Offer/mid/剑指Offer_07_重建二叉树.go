package mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}

	value := preorder[0]
	pivo := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == value {
			pivo = i
			break
		}
	}

	root := &TreeNode{
		Val:   value,
		Left:  buildHelper(1, pivo, 0, pivo-1, &preorder, &inorder),
		Right: buildHelper(pivo+1, len(preorder)-1, pivo+1, len(inorder)-1, &preorder, &inorder),
	}
	return root
}

func buildHelper(pS int, pE int, iS int, iE int, preorder *[]int, inorder *[]int) *TreeNode {
	if pS > pE || iS > iE {
		return nil
	}
	value := (*preorder)[pS]
	pivo := 0
	for i := iS; i <= iE; i++ {
		if (*inorder)[i] == value {
			pivo = i
			break
		}
	}

	root := &TreeNode{
		Val:   value,
		Left:  buildHelper(pS+1, pivo-iS+pS, iS, pivo-1, preorder, inorder),
		Right: buildHelper(pivo-iS+pS+1, pE, pivo+1, iE, preorder, inorder),
	}
	return root
}
