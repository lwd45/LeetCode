package mid

func findChampion(n int, edges [][]int) int {
	inDegree := make([]int, n)
	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	ans := -1
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 && ans == -1 {
			ans = i
		} else if inDegree[i] == 0 {
			return -1
		}
	}
	return ans
}
