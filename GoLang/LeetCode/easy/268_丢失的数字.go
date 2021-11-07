package easy

func missingNumber(nums []int) int {
	sum := (1 + len(nums)) * len(nums) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
