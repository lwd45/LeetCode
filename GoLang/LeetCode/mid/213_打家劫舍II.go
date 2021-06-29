package mid

/*
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(solution(&nums, 0, len(nums)-2), solution(&nums, 1, len(nums)-1))
}

func solution(nums *[]int, start, end int) int {
	first, second := 0, (*nums)[start]
	for i := start + 1; i <= end; i++ {
		temp := max(first+(*nums)[i], second)
		first = second
		second = temp
	}
	return second
}

func max(nums ...int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}
