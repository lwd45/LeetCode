package easy

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	total := 0
	for i := 0; i < len(nums); i++ {
		if total == (sum - total - nums[i]) {
			return i
		} else {
			total += nums[i]
		}
	}
	return -1
}
