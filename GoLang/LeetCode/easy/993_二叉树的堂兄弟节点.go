package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func IsCousins(root *TreeNode, x int, y int) bool {
//	if root.Val == x || root.Val == y {
//		return false
//	}
//
//	nodes, contains, oneNodeContains := make([]*TreeNode, 0, 1), false, false
//	nodes = append(nodes, root)
//
//	for len(nodes) > 0 {
//		nextNodes := make([]*TreeNode, 0, len(nodes)*2)
//		for len(nodes) > 0 {
//			poll := nodes[0]
//			nodes = nodes[1:]
//
//			if poll.Left != nil {
//				nextNodes = append(nextNodes, poll.Left)
//				if contains && (poll.Left.Val == x || poll.Left.Val == y) {
//					return true
//				} else if poll.Left.Val == x || poll.Left.Val == y {
//					contains = true
//					oneNodeContains = true
//				}
//			}
//			if poll.Right != nil {
//				nextNodes = append(nextNodes, poll.Right)
//				if oneNodeContains && (poll.Right.Val == x || poll.Right.Val == y) {
//					return false
//				} else if (poll.Right.Val == x || poll.Right.Val == y) && contains {
//					return true
//				} else if poll.Right.Val == x || poll.Right.Val == y {
//					contains = true
//				}
//			}
//
//			if oneNodeContains {
//				oneNodeContains = !oneNodeContains
//			}
//		}
//
//		if contains {
//			return false
//		}
//		nodes = nextNodes
//	}
//	return false
//}

/*
func isCousins(root *TreeNode, x, y int) bool {
    var xParent, yParent *TreeNode
    var xDepth, yDepth int
    var xFound, yFound bool

    var dfs func(node, parent *TreeNode, depth int)
    dfs = func(node, parent *TreeNode, depth int) {
        if node == nil {
            return
        }

        if node.Val == x {
            xParent, xDepth, xFound = parent, depth, true
        } else if node.Val == y {
            yParent, yDepth, yFound = parent, depth, true
        }

        // 如果两个节点都找到了，就可以提前退出遍历
        // 即使不提前退出，对最坏情况下的时间复杂度也不会有影响
        if xFound && yFound {
            return
        }
        dfs(node.Left, node, depth+1)
        if xFound && yFound {
            return
        }
        dfs(node.Right, node, depth+1)
    }
    dfs(root, nil, 0)
    return xDepth == yDepth && xParent != yParent
}
*/

var xParent, yParent *TreeNode
var xFind, yFind bool
var xDeep, yDeep int

func isCousins(root *TreeNode, x, y int) bool {
	xFind, yFind = false, false
	dfs(root, nil, x, y, 1)
	return xFind && yFind && xDeep == yDeep && xParent != yParent
}

func dfs(root, parent *TreeNode, x, y, deep int) {
	if (xFind && yFind) || root == nil {
		return
	}

	if root.Val == x {
		xFind, xParent, xDeep = true, parent, deep
	} else if root.Val == y {
		yFind, yParent, yDeep = true, parent, deep
	}
	dfs(root.Left, root, x, y, deep+1)
	dfs(root.Right, root, x, y, deep+1)
}
