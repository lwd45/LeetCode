package mid

func sumDigitDifferences(nums []int) int64 {
	ans := 0
	for nums[0] > 0 {
		cnt := make([]int, 10)
		for i, num := range nums {
			cnt[num%10]++
			nums[i] /= 10
		}

		for _, c := range cnt {
			ans += c * (len(nums) - c)
		}
	}
	return int64(ans) / 2
}
