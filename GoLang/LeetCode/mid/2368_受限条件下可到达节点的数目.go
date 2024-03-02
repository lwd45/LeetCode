package mid

func reachableNodes(n int, edges [][]int, restricted []int) int {
	eMap := make(map[int][]int, len(edges))
	restrictedMap := make(map[int]bool)
	visited := make(map[int]bool)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		eMap[from] = append(eMap[from], to)
		eMap[to] = append(eMap[to], from)
	}
	for _, node := range restricted {
		restrictedMap[node] = true
	}

	ans := 1
	visited[0] = true
	nexts := eMap[0]

	for len(nexts) > 0 {
		var temp []int
		for _, node := range nexts {
			if _, ok := restrictedMap[node]; ok { // 是限制节点
				continue
			}
			if _, ok := visited[node]; ok {
				continue
			}

			ans++
			visited[node] = true
			temp = append(temp, eMap[node]...)
		}
		nexts = temp
	}
	return ans
}
