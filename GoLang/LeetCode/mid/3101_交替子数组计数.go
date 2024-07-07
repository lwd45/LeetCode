package mid

func countAlternatingSubarrays(nums []int) int64 {
	ans := 0

	pre := -1
	cnt := 0
	for _, num := range nums {
		if pre != num {
			cnt++
		} else {
			cnt = 1
		}

		pre = num
		ans += cnt
	}
	return int64(ans)
}
