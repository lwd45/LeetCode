package mid

//func LongestSubarray(nums []int) int {
//	left, right, length := 0, 0, len(nums)
//	isContainZero, zeroIndex, hasZeroInNums := false, -1, false
//	maxLen := 0
//
//	for ; left < length && nums[left] != 1; left++ {
//		hasZeroInNums = true
//	}
//	if left >= length {
//		return 0
//	}
//
//	for right = left; right < length; right++ {
//		if nums[right] == 0 && !isContainZero {
//			isContainZero, zeroIndex, hasZeroInNums = true, right, true
//		} else if nums[right] == 0 {
//			maxLen = max(maxLen, right-left-1)
//			left = zeroIndex + 1
//			zeroIndex = right
//		}
//	}
//
//	if !hasZeroInNums {
//		return length - 1
//	}
//	maxLen = max(maxLen, right-left-1)
//	if !isContainZero {
//		return maxLen + 1
//	}
//	return maxLen
//}
//
//func max(num1, num2 int) int {
//	if num1 > num2 {
//		return num1
//	}
//	return num2
//}

//func longestSubarray(nums []int) int {
//	left, right, deleteCount := 0, 0, 1
//	for ; right < len(nums); right++ {
//		if nums[right] == 0 {
//			deleteCount--
//		}
//		if deleteCount < 0 {
//			if nums[left] == 0 {
//				deleteCount++
//			}
//			left++
//		}
//	}
//	return right - left - 1
//}

//func longestSubarray(nums []int) int {
//	left, right, deleteCount := 0, 0, 0
//
//	for ; right < len(nums); right++ {
//		if nums[right] == 0 {
//			deleteCount++
//		}
//		if deleteCount > 1 {
//			if nums[left] == 0 {
//				deleteCount--
//			}
//			left++
//		}
//	}
//	return right - left - 1
//}

func longestSubarray(nums []int) int {
	left, right, deleteCount := 0, 0, 1

	for ; right < len(nums); {
		if nums[right] == 0 {
			deleteCount--
			if deleteCount < 0 {
				for deleteCount < 0 && right < len(nums) {
					if nums[left] == 0 {
						deleteCount++
					}
					left++
					right++
					if right < len(nums) && nums[right] == 0 {
						deleteCount--
					}
				}
			} else {
				right++
			}
		} else {
			right++
		}
	}
	return right - left - 1
}
