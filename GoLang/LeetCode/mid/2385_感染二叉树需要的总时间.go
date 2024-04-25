package mid

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)
	dfs_2385(root, graph)

	visited := make(map[int]bool)
	visited[start] = true

	ans := 0
	nodes := graph[start]
	for len(nodes) != 0 {
		ans++
		newNodes := make([]int, 0, len(nodes))

		for _, node := range nodes {
			visited[node] = true
			for _, ns := range graph[node] {
				if !visited[ns] {
					newNodes = append(newNodes, ns)
				}
			}
		}

		nodes = newNodes
	}
	return ans
}

func dfs_2385(root *TreeNode, graph map[int][]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		graph[root.Val] = append(graph[root.Val], root.Left.Val)
		graph[root.Left.Val] = append(graph[root.Left.Val], root.Val)
		dfs_2385(root.Left, graph)
	}
	if root.Right != nil {
		graph[root.Val] = append(graph[root.Val], root.Right.Val)
		graph[root.Right.Val] = append(graph[root.Right.Val], root.Val)
		dfs_2385(root.Right, graph)
	}
}
