package easy

func minMoves(nums []int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	ans := 0
	for _, num := range nums {
		ans += num - min
	}
	return ans
}
