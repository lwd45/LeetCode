package hard

func kInversePairs(n int, k int) int {
	mod := 1000000007
	dp := make([][]int, n+1)
	for idx := range dp {
		dp[idx] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k && j <= (i-1)*i/2; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else {
				val := 0
				if j-i >= 0 {
					val = (dp[i-1][j] - dp[i-1][j-i] + mod) % mod
				} else {
					val = dp[i-1][j]
				}
				dp[i][j] = (dp[i][j-1] + val) % mod
			}
		}
	}
	return dp[n][k]
}
