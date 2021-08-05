package mid

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right, max := 0, 0, 0
	cnt := make(map[byte]int, len(s))
	for ; right < len(s); right++ {
		cnt[s[right]]++
		for cnt[s[right]] > 1 {
			cnt[s[left]]--
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}
