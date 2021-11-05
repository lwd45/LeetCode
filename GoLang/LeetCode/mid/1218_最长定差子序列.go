package mid

func longestSubsequence(arr []int, difference int) int {
	ans := 1
	cnt := make(map[int]int, len(arr))
	for _, key := range arr {
		if value, ok := cnt[key-difference]; ok {
			cnt[key] = value + 1
			if cnt[key] > ans {
				ans = cnt[key]
			}
		} else {
			cnt[key] = 1
		}
	}
	return ans
}
