package mid

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i, _ := range s {
		dp[i] = make([]int, len(s))
	}

	for j := 0; j < len(s); j++ {
		dp[j][j] = 1
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max_516(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max_516(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
