package mid

func minIncrements(n int, cost []int) int {
	ans := 0
	for idx := n - 2; idx > 0; idx -= 2 {
		ans += abs(cost[idx], cost[idx+1])
		cost[idx/2] = cost[idx/2] + max(cost[idx], cost[idx+1])
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
