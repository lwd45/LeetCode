package mid

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	e := make([][]int, n)
	inDegree := make([]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		e[from] = append(e[from], to)
		inDegree[to]++
	}

	var q []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	anc := make([]map[int]bool, n) // 存储每个节点祖先的辅助数组
	for i := 0; i < n; i++ {
		anc[i] = map[int]bool{}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range e[u] {
			anc[v][u] = true
			for key, val := range anc[u] {
				if val {
					anc[v][key] = true
				}
			}
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		for key, val := range anc[i] {
			if val {
				res[i] = append(res[i], key)
			}
		}
		sort.Ints(res[i])
	}
	return res
}

//func getAncestors(n int, edges [][]int) [][]int {
//	g := make([][]int, n)
//	for _, edge := range edges {
//		g[edge[0]] = append(g[edge[0]], edge[1])
//	}
//
//	ans := make([][]int, n)
//	start := 0
//
//	vis := make([]int, n)
//	var dfs func(x int)
//	dfs = func(x int) {
//		vis[x] = start + 1
//		for _, e := range g[x] {
//			if vis[e] != start+1 {
//				ans[e] = append(ans[e], start)
//				dfs(e)
//			}
//		}
//	}
//
//	for ; start < n; start++ {
//		dfs(start)
//	}
//	return ans
//}

//func getAncestors(n int, edges [][]int) [][]int {
//	g := make([][]int, n)
//	for _, edge := range edges {
//		g[edge[1]] = append(g[edge[1]], edge[0])
//	}
//
//	vis := make([]bool, n)
//	var dfs func(x int)
//	dfs = func(x int) {
//		vis[x] = true
//		for _, e := range g[x] {
//			if !vis[e] {
//				dfs(e)
//			}
//		}
//	}
//
//	ans := make([][]int, n)
//	for i := 0; i < n; i++ {
//		vis = make([]bool, n)
//		dfs(i)
//		vis[i] = false
//
//		for idx, v := range vis {
//			if v {
//				ans[i] = append(ans[i], idx)
//			}
//		}
//	}
//	return ans
//}
