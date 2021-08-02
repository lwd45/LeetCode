package easy

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
