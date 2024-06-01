package mid

//func maximumLength(s string) int {
//	ans := -1
//
//	cnt := make(map[string]int)
//	for i, c := range s {
//		cnt[string(c)]++
//		if cnt[string(c)] >= 3 {
//			ans = max(ans, 1)
//		}
//
//		j := i
//		for ; j > 0 && s[i] == s[j-1]; j-- {
//			str := s[j-1 : i+1]
//			cnt[str]++
//			if cnt[str] >= 3 {
//				ans = max(ans, len(str))
//			}
//		}
//	}
//	return ans
//}

func maximumLength(s string) int {
	lens := len(s)

	cnt := 0
	chs := make([][]int, 26)
	for i := 0; i < lens; i++ {
		cnt++
		if i+1 == lens || s[i] != s[i+1] {
			c := s[i] - 'a'
			chs[c] = append(chs[c], cnt)
			cnt = 0

			for j := len(chs[c]) - 1; j > 0; j-- {
				if chs[c][j] > chs[c][j-1] {
					t := chs[c][j]
					chs[c][j] = chs[c][j-1]
					chs[c][j-1] = t
				} else {
					break
				}
			}
			if len(chs[c]) > 3 {
				chs[c] = chs[c][:3]
			}
		}
	}

	ans := -1
	for _, ch := range chs {
		if len(ch) > 0 && ch[0] >= 3 {
			ans = max(ans, ch[0]-2)
		}
		if len(ch) > 1 && ch[0] >= 2 {
			ans = max(ans, min(ch[0]-1, ch[1]))
		}
		if len(ch) > 2 {
			ans = max(ans, ch[2])
		}
	}
	return ans
}
