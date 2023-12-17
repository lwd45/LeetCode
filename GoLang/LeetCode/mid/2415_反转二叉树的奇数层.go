package mid

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func reverseOddLevels(root *TreeNode) *TreeNode {
//	if root == nil {
//		return root
//	}
//
//	nodes := make([]*TreeNode, 0, 1)
//	nodes = append(nodes, root)
//	level := 0
//	for len(nodes) > 0 {
//		if level%2 == 1 { // 置换val
//			for i, j := 0, len(nodes)-1; i < j; {
//				t := nodes[i].Val
//				nodes[i].Val = nodes[j].Val
//				nodes[j].Val = t
//				i++
//				j--
//			}
//		}
//
//		temp := make([]*TreeNode, 0, 2*len(nodes))
//		for _, node := range nodes {
//			if node.Left != nil {
//				temp = append(temp, node.Left, node.Right)
//			}
//		}
//		nodes = temp
//		level++
//	}
//	return root
//}

func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs_2415(root.Left, root.Right, true)
	return root
}

func dfs_2415(left *TreeNode, right *TreeNode, isOdd bool) {
	if left == nil {
		return
	}

	if isOdd {
		left.Val, right.Val = right.Val, left.Val
	}

	dfs_2415(left.Left, right.Right, !isOdd)
	dfs_2415(left.Right, right.Left, !isOdd)
}
