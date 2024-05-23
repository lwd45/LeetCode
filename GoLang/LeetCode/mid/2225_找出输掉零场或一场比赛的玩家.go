package mid

import "sort"

func findWinners(matches [][]int) [][]int {

	winMap, loseMap := make(map[int]int), make(map[int]int)
	for _, match := range matches {
		winner := match[0]
		loser := match[1]

		winMap[winner]++
		loseMap[loser]++
	}
	ans := make([][]int, 2)

	for key, _ := range winMap {
		if _, ok := loseMap[key]; !ok {
			ans[0] = append(ans[0], key)
		}
	}
	for key, v := range loseMap {
		if v == 1 {
			ans[1] = append(ans[1], key)
		}
	}
	sort.Slice(ans[0], func(i, j int) bool {
		return ans[0][i] < ans[0][j]
	})
	sort.Slice(ans[1], func(i, j int) bool {
		return ans[1][i] < ans[1][j]
	})
	return ans
}
