package mid

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	profitMap := make(map[int]int, n)
	for i := 0; i < n; i++ {
		profitMap[difficulty[i]] = max(profitMap[difficulty[i]], profit[i])
	}

	sort.Ints(difficulty)
	sort.Ints(worker)

	maxProfit, idx := profitMap[difficulty[0]], 0

	ans := 0
	for i := 0; i < len(worker); i++ {
		if worker[i] < difficulty[idx] {
			continue
		}

		for idx < n && worker[i] >= difficulty[idx] {
			maxProfit = max(maxProfit, profitMap[difficulty[idx]])
			idx++
		}
		idx--

		ans += maxProfit
	}
	return ans
}
