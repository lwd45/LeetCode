package mid

func validPartition(nums []int) bool {
	if len(nums) == 2 {
		return nums[1] == nums[0]
	}
	if len(nums) == 3 {
		return (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[0]+1 == nums[1] && nums[1]+1 == nums[2])
	}

	dp := make([]bool, len(nums))
	dp[1] = nums[1] == nums[0]
	dp[2] = (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[0]+1 == nums[1] && nums[1]+1 == nums[2])

	for i := 3; i < len(nums); i++ {
		dp[i] = dp[i-2] && (nums[i] == nums[i-1])                                        // 2个元素相等
		dp[i] = dp[i] || dp[i-3] && (nums[i] == nums[i-1] && nums[i-1] == nums[i-2])     // 3个元素相等
		dp[i] = dp[i] || dp[i-3] && (nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1) // 3个元素相差1
	}
	return dp[len(nums)-1]
}
