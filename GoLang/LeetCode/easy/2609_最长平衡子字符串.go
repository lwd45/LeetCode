package easy

func findTheLongestBalancedSubstring(s string) int {
	flag := 0 // 0表示统计0， 1表示统计1
	cnt := 0

	res, t := 0, 0
	for _, c := range s {
		if c == '0' {
			if flag == 0 {
				cnt++
			} else {
				cnt = 1
			}

			flag = 0
			t = 0
		} else {
			if cnt > 0 {
				t++
				if 2*t > res {
					res = 2 * t
				}
			}
			cnt--
			flag = 1
		}
	}
	return res
}
