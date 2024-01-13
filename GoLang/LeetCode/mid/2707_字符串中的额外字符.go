package mid

import "math"

func minExtraChar(s string, dictionary []string) int {
	mp := make(map[string]struct{}, len(dictionary))
	for _, v := range dictionary {
		mp[v] = struct{}{}
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if _, ok := mp[s[j:i]]; ok {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}
