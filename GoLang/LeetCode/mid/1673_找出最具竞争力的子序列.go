package mid

func mostCompetitive(nums []int, k int) []int {
	var ans []int

	for i := 0; i < len(nums); i++ {
		for len(ans) > 0 && nums[i] < ans[len(ans)-1] && len(nums)-i >= k-len(ans)+1 {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, nums[i])
		}
	}
	return ans
}
