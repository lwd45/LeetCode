package main

func main() {
	maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}})
}

func maximumDetonation(bombs [][]int) int {
	visited := make(map[int]bool, len(bombs))
	countMap := make(map[int]int)
	for i := 0; i < len(bombs); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		countMap[i]++

		dfs_2101(bombs, i, i, visited, countMap)
	}

	ans := 1
	for _, v := range countMap {
		ans = max(ans, v)
	}
	return ans
}

func dfs_2101(bombs [][]int, start, idx int, visited map[int]bool, countMap map[int]int) {
	for i := 0; i < len(bombs); i++ {
		if visited[i] {
			continue
		}

		dx := bombs[idx][0] - bombs[i][0]
		dy := bombs[idx][1] - bombs[i][1]
		r := bombs[idx][2]

		if dx*dx+dy*dy <= r*r {
			visited[i] = true
			countMap[start]++
			dfs_2101(bombs, start, i, visited, countMap)
		}
	}
}
