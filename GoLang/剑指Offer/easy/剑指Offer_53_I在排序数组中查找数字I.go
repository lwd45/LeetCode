package easy

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1

	for left <= right && nums[left] != target {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid
			break
		}
	}

	if left >= len(nums) || nums[left] != target {
		return 0
	}
	l, r := left, left
	for l >= 0 && nums[l] == target {
		l--
	}
	for r < len(nums) && nums[r] == target {
		r++
	}
	return r - l - 1
}
