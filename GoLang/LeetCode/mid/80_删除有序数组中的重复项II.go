package mid

func removeDuplicates(nums []int) int {
	index, flag := 0, false
	for i := 1; i < len(nums); i++ {
		if flag && nums[i] == nums[index] {
			continue
		} else if nums[i] == nums[index] {
			flag = true
			index++
			nums[index] = nums[i]
		} else {
			flag = false
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
