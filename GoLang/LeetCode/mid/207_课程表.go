package mid

//func canFinish(numCourses int, prerequisites [][]int) bool {
//	if len(prerequisites) == 0 {
//		return true
//	}
//
//	edges := make([][]int, numCourses)
//	for _, prerequisite := range prerequisites {
//		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
//	}
//
//	valid := true
//	visit := make([]int, numCourses)
//	for i := 0; i < numCourses; i++ {
//		if visit[i] == 0 {
//			dfss(i, edges, visit, &valid)
//		}
//	}
//	return valid
//}
//
//func dfss(i int, edges [][]int, visit []int, valid *bool) {
//	if !*valid {
//		return
//	}
//
//	visit[i] = 1 // 搜索中
//	for _, v := range edges[i] {
//		if visit[v] == 0 { // 未被搜索
//			dfss(v, edges, visit, valid)
//		} else if visit[v] == 1 {
//			*valid = false
//			return
//		}
//	}
//	visit[i] = 2 // 搜索完
//}

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	var result []int
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		result = append(result, v)
		for _, edge := range edges[v] {
			indeg[edge]--
			if indeg[edge] == 0 {
				q = append(q, edge)
			}
		}
	}
	return len(result) == numCourses
}
