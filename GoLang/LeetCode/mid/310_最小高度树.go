package mid

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	paths := make(map[int][]int, n)
	degree := make(map[int]int, n)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		if _, ok := paths[node1]; !ok {
			paths[node1] = []int{}
		}
		if _, ok := paths[node2]; !ok {
			paths[node2] = []int{}
		}
		paths[node1] = append(paths[node1], node2)
		paths[node2] = append(paths[node2], node1)
		degree[node1]++
		degree[node2]++
	}

	var queue []int
	for key, val := range degree {
		if val == 1 {
			queue = append(queue, key)
		}
	}

	var res []int
	for len(queue) != 0 {
		res = []int{}

		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			res = append(res, node)
			for _, neighbourNode := range paths[node] {
				degree[neighbourNode]--
				if degree[neighbourNode] == 1 {
					queue = append(queue, neighbourNode)
				}
			}
		}
	}
	return res
}
