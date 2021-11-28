package mid

func findAnagrams(s string, p string) (ans []int) {
	len1, len2 := len(s), len(p)
	if len2 > len1 {
		return
	}

	cntP := make(map[int32]int)
	for _, c := range p {
		cntP[c] += 1
	}

	start := 0
	cnt := make(map[int32]int)
	for end, c := range s {
		cnt[c] += 1
		for ; cnt[c] > cntP[c]; {
			cnt[int32(s[start])] -= 1
			start++
		}
		if end-start+1 == len2 {
			ans = append(ans, start)
			cnt[int32(s[start])] -= 1
			start++
		}
	}
	return
}
