package mid

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := root
	for {
		if ans.Val > p.Val && ans.Val > q.Val {
			ans = ans.Left
		} else if ans.Val < p.Val && ans.Val < q.Val {
			ans = ans.Right
		} else {
			break
		}
	}
	return ans
}

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	pPath := getPath(root, p)
//	qPath := getPath(root, q)
//
//	for pPath[len(pPath)-1].Val != qPath[len(qPath)-1].Val{
//		if len(pPath) > len(qPath) {
//			pPath = pPath[:len(pPath)-1]
//		} else if len(pPath) < len(qPath) {
//			qPath = qPath[:len(qPath)-1]
//		} else if pPath[len(pPath)-1].Val != qPath[len(qPath)-1].Val {
//			pPath = pPath[:len(pPath)-1]
//			qPath = qPath[:len(qPath)-1]
//		}
//	}
//	return pPath[len(pPath)-1]
//}
//
//func getPath(root, target *TreeNode) []*TreeNode {
//	var paths []*TreeNode
//	node := root
//	for node.Val != target.Val {
//		paths = append(paths, node)
//		if node.Val > target.Val {
//			node = node.Left
//		} else {
//			node = node.Right
//		}
//	}
//	paths = append(paths, node)
//	return paths
//}
