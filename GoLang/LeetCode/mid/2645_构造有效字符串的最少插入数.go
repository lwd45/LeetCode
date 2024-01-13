package mid

func addMinimum(word string) int {
	n := len(word)
	dp := make([]int, n)

	dp[0] = 2
	for i := 1; i < n; i++{
		if word[i] > word[i-1]{
			dp[i] = dp[i-1]-1
		}else{
			dp[i] = dp[i-1]+2
		}
	}
	return dp[n-1]
}