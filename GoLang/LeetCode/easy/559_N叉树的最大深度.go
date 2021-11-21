package easy

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    maxChildDepth := 0
    for _, child := range root.Children {
        if childDepth := maxDepth(child); childDepth > maxChildDepth {
            maxChildDepth = childDepth
        }
    }
    return maxChildDepth + 1
}
*/

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) (ans int) {
	if root == nil {
		return 0
	}
	maxDeepValue := 0
	for _, child := range root.Children {
		childDeep := maxDepth(child)
		if maxDeepValue < childDeep {
			maxDeepValue = childDeep
		}
	}
	return maxDeepValue + 1
}

//func maxDepth(root *Node) (ans int) {
//	if root == nil {
//		return 0
//	}
//	list := []*Node{root}
//	for len(list) > 0 {
//		ans++
//		temp := list
//		list = nil
//		for _, node := range temp {
//			list = append(list, node.Children...)
//		}
//	}
//	return ans
//}
