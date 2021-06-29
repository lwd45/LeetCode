package mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorderTraversal(root *TreeNode) []int {
//	var ans []int
//	inorder(root, &ans)
//	return ans
//}
//
//func inorder(root *TreeNode, ans *[]int) {
//	if root == nil {
//		return
//	}
//	inorder(root.Left, ans)
//	*ans = append(*ans, root.Val)
//	inorder(root.Right, ans)
//}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var nodeList []*TreeNode
	var ans []int
	for root != nil || len(nodeList) > 0 {
		for root != nil {
			nodeList = append(nodeList, root)
			root = root.Left
		}
		root = nodeList[len(nodeList)-1]
		ans = append(ans, root.Val)
		nodeList = nodeList[:len(nodeList)-1]
		root = root.Right
	}
	return ans
}
