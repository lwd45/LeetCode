package mid

func minCost(nums []int, x int) int64 {
	n := len(nums)
	ans := sum(nums)

	f := make([]int, n)
	copy(f, nums)

	for k := 1; k < n; k++ {
		for i := 0; i < n; i++ {
			f[i] = min(f[i], nums[(i+k)%n])
		}
		ans = min(ans, sum(f)+x*k)
	}
	return int64(ans)
}

func sum(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}
