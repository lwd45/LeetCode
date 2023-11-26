package mid

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt32
		}
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		mp[from][to], mp[to][from] = weight, weight
	}

	for k := 0; k < n; k++ {
		mp[k][k] = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				mp[i][j] = min(mp[i][j], mp[i][k]+mp[k][j])
			}
		}
	}

	ans := []int{math.MaxInt32, -1}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if mp[i][j] <= distanceThreshold {
				cnt++
			}
		}

		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}
	return ans[1]
}
