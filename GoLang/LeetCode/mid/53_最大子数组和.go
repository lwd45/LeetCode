package mid

import "math"

func maxSubArray(nums []int) int {
	ans, sum := math.MinInt64, 0
	for _, num := range nums {
		sum += num
		ans = max(ans, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}
