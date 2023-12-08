package mid

func maxTaxiEarnings(n int, rides [][]int) int64 {
	rideMap := make(map[int][][]int)
	for _, ride := range rides {
		_, end, _ := ride[0], ride[1], ride[2]
		rideMap[end] = append(rideMap[end], ride)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for _, ride := range rideMap[i] {
			start, end, tip := ride[0], ride[1], ride[2]
			dp[i] = max(dp[i], dp[start]+end-start+tip)
		}
	}
	return int64(dp[n])
}
