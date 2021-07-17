package easy

func maxSubArray(nums []int) int {
	max, sum := nums[0], 0

	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
