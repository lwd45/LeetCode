package mid

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	edges := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}

	valid := true
	visit := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if visit[i] == 0 {
			dfss(i, edges, visit, &valid)
		}
	}
	return valid
}

func dfss(i int, edges [][]int, visit []int, valid *bool) {
	if !*valid {
		return
	}

	visit[i] = 1 // 搜索中
	for _, v := range edges[i] {
		if visit[v] == 0 { // 未被搜索
			dfss(v, edges, visit, valid)
		} else if visit[v] == 1 {
			*valid = false
			return
		}
	}
	visit[i] = 2 // 搜索完
}
