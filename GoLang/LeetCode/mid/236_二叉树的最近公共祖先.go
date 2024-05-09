package mid

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil || root == p || root == q {
//		return root
//	}
//
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//
//	if left == nil && right == nil {
//		return nil
//	}
//	if left == nil {
//		return right
//	}
//	if right == nil {
//		return left
//	}
//	return root
//}

//	func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//		parent := map[int]*TreeNode{}
//		visited := map[int]bool{}
//
//		dfs_236(root, parent)
//		for p != nil {
//			visited[p.Val] = true
//			p = parent[p.Val]
//		}
//		for q != nil {
//			if visited[q.Val] {
//				return q
//			}
//			q = parent[q.Val]
//		}
//		return nil
//	}
//
//	func dfs_236(root *TreeNode, parent map[int]*TreeNode) {
//		if root == nil {
//			return
//		}
//
//		if root.Left != nil {
//			parent[root.Left.Val] = root
//			dfs_236(root.Left, parent)
//		}
//		if root.Right != nil {
//			parent[root.Right.Val] = root
//			dfs_236(root.Right, parent)
//		}
//	}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	return root
}
