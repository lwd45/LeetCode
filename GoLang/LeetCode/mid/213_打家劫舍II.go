package mid

/*
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/
//func rob(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	return max(solution(&nums, 0, len(nums)-2), solution(&nums, 1, len(nums)-1))
//}
//
//func solution(nums *[]int, start, end int) int {
//	first, second := 0, (*nums)[start]
//	for i := start + 1; i <= end; i++ {
//		temp := max(first+(*nums)[i], second)
//		first = second
//		second = temp
//	}
//	return second
//}
//
//func max(nums ...int) int {
//	ret := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > ret {
//			ret = nums[i]
//		}
//	}
//	return ret
//}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
