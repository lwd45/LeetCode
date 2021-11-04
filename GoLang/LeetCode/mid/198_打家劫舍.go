package mid

func robs(nums []int) int {
	length := len(nums)
	dp := make([][2]int, length)
	dp[0][1] = nums[0] //dp[i][0]表示不偷第i家	dp[i][1]表示偷第i家
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = max(dp[i-1][0]+nums[i], dp[i-1][1])
	}
	return max(dp[length-1][0], dp[length-1][1])
}
