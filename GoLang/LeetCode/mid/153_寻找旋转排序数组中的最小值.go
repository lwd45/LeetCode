package mid

func findMin(nums []int) int {
	length := len(nums)
	if nums[0] <= nums[length-1] {
		return nums[0]
	}

	L, R, Mid := 0, length-1, 0
	for L < R && nums[L] > nums[R] {
		Mid = (L + R) >> 1
		if nums[Mid] >= nums[L] {
			L = Mid + 1
		} else {
			R = Mid
		}
	}
	return nums[L]
}
