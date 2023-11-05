package mid

func findRepeatedDnaSequences(s string) []string {
	var ans []string
	cnt := make(map[string]int, 0)
	for idx := 0; idx <= len(s)-10; idx++ {
		s := s[idx : idx+10]
		cnt[s]++
		if cnt[s] == 2 {
			ans = append(ans, s)
		}
	}
	return ans
}
