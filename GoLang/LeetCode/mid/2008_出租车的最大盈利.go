package mid

import "sort"

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	m := len(rides)
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		j := sort.Search(i+1, func(k int) bool {
			return rides[k][1] > rides[i][0]
		})
		dp[i] = max(dp[i-1], dp[j]+rides[j][1]-rides[j][0]+rides[j][2])
	}
	return int64(dp[m])
}

//func maxTaxiEarnings(n int, rides [][]int) int64 {
//	rideMap := make(map[int][][]int)
//	for _, ride := range rides {
//		_, end, _ := ride[0], ride[1], ride[2]
//		rideMap[end] = append(rideMap[end], ride)
//	}
//
//	dp := make([]int, n+1)
//	for i := 1; i <= n; i++ {
//		dp[i] = dp[i-1]
//		for _, ride := range rideMap[i] {
//			start, end, tip := ride[0], ride[1], ride[2]
//			dp[i] = max(dp[i], dp[start]+end-start+tip)
//		}
//	}
//	return int64(dp[n])
//}
