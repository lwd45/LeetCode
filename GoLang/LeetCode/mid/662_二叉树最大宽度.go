package mid

//type pair struct {
//	root  *TreeNode
//	index int
//}

//func widthOfBinaryTree(root *TreeNode) int {
//	pairList := []pair{{
//		root:  root,
//		index: 1,
//	}}
//
//	ans := 1
//	for len(pairList) > 0 {
//		ans = max(ans, pairList[len(pairList)-1].index-pairList[0].index+1)
//		temp := pairList
//		pairList = nil
//		for _, p := range temp {
//			if p.root.Left != nil {
//				pairList = append(pairList, pair{
//					root:  p.root.Left,
//					index: p.index * 2,
//				})
//			}
//
//			if p.root.Right != nil {
//				pairList = append(pairList, pair{
//					root:  p.root.Right,
//					index: p.index*2 + 1,
//				})
//			}
//		}
//	}
//
//	return ans
//}

func widthOfBinaryTree(root *TreeNode) int {
	deepIdxMap := map[int]int{}
	return dfs(root, 1, 1, deepIdxMap)
}

func dfs(root *TreeNode, deep, index int, deepIdxMap map[int]int) int {
	if root == nil {
		return 0
	}

	if _, ok := deepIdxMap[deep]; !ok {
		deepIdxMap[deep] = index
	}

	return max(index-deepIdxMap[deep]+1,
		max(dfs(root.Left, deep+1, index*2, deepIdxMap), dfs(root.Right, deep+1, index*2+1, deepIdxMap)))
}
