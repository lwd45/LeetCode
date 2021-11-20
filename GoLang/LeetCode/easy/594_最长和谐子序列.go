package easy

func findLHS(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num] += 1
	}
	ans := 0
	for _, num := range nums {
		if _, ok := cnt[num+1]; ok {
			ans = max(ans, cnt[num]+cnt[num+1])
		}
	}
	return ans
}

func max(nums ...int) int {
	ans := nums[0]
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans
}
