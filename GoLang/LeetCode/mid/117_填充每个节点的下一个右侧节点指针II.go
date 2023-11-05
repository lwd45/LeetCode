package mid

/**
 * Definition for a Node.
 *
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	nodes := make([]*Node, 0, 1)
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		newNodes := make([]*Node, 0, 2*len(nodes))
		for i := 0; i < len(nodes); i++ {
			if i != len(nodes)-1 {
				nodes[i].Next = nodes[i+1]
			}

			if nodes[i].Left != nil {
				newNodes = append(newNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				newNodes = append(newNodes, nodes[i].Right)
			}
		}
		nodes = newNodes
	}
	return root
}
