package easy

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}
	for i, c := range s {
		if cnt[c-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
