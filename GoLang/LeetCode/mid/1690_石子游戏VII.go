package mid

func stoneGameVII(stones []int) int {
	n := len(stones)

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
	}

	dp := make([][]int, n) // dp[i][j] 表示剩下的数据中，先手比后手的差值
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(sum[j]-sum[i]-dp[i][j-1], sum[j+1]-sum[i+1]-dp[i+1][j])
		}
	}
	return dp[0][n-1]
}
