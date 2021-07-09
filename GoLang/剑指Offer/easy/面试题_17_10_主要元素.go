package easy

func majorityElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	num, count := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count++
		} else {
			count--
			if count < 0 {
				num = nums[i]
				count = 0
			}
		}
	}
	cnt := 0
	for _, v := range nums {
		if v == num {
			cnt++
		}
	}
	if cnt <= len(nums)/2 {
		return -1
	}
	return num
}
