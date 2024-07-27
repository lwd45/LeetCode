package mid

//func getSmallestString(s string, k int) string {
//	ans := make([]int32, 0, len(s))
//
//	for i, c := range s {
//		if k == 0 {
//			return string(ans) + s[i:]
//		}
//
//		if c == 'a' {
//			ans = append(ans, 'a')
//			continue
//		} else {
//			minToA := 26 - int(c-'a')
//			if int(c-'a') < minToA {
//				minToA = int(c - 'a')
//			}
//
//			if minToA <= k {
//				ans = append(ans, 'a')
//				k -= minToA
//			} else {
//				ans = append(ans, c-int32(k))
//				k = 0
//			}
//		}
//	}
//
//	return string(ans)
//}

func getSmallestString(s string, k int) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		dis := min(int(runes[i]-'a'), int('z'-runes[i]+1))
		if dis <= k {
			runes[i] = 'a'
			k -= dis
		} else {
			runes[i] = rune(int(runes[i]) - k)
			break
		}
	}
	return string(runes)
}
