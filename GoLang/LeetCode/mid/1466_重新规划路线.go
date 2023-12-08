package mid

func minReorder(n int, connections [][]int) int {
	mF := make(map[int][]int, len(connections))
	mT := make(map[int][]int, len(connections))
	for _, connection := range connections {
		from, to := connection[0], connection[1]
		mF[from] = append(mF[from], to)
		mT[to] = append(mT[to], from)
	}

	ans := 0
	visited := make(map[int]bool, len(connections))

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for _, from := range mT[i] {
			if !visited[from] {
				dfs(from)
			}
		}
		for _, to := range mF[i] {
			if !visited[to] {
				ans += 1
				dfs(to)
			}
		}
	}
	dfs(0)
	return ans
}
