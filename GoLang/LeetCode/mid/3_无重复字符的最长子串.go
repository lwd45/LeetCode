package mid

/*
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	max, cSet := 1, map[uint8]struct{}{}
//	for L, R := 0, 0; R < len(s); R++ {
//		for _, ok := cSet[s[R]]; ok; _, ok = cSet[s[R]] {
//			delete(cSet, s[L])
//			L++
//		}
//		cSet[s[R]] = struct{}{}
//		if max < R-L+1 {
//			max = R - L + 1
//		}
//	}
//	return max
//}

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//
//	start := 0
//	end := 0
//	var appeared [256]int
//	for i := 0; i < len(appeared); i++ {
//		appeared[i] = -1
//	}
//	max := 0
//	for end < len(s) {
//		if idx := appeared[s[end] - 'a']; idx != -1 && idx >= start {
//			start = idx + 1
//		}
//		tmp := end - start + 1
//		if tmp > max {
//			max = tmp
//		}
//		appeared[s[end]-'a'] = end
//		end++
//	}
//	return max
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ans := 0
	cnt := [256]int32{}
	left, right, m := 0, 0, len(s)

	for right < m {
		idx := s[right] - 'a'
		cnt[idx]++

		if cnt[idx] > 1 {
			for left < right {
				i2 := s[left] - 'a'
				cnt[i2]--
				left++
				if cnt[idx] == 1 {
					break
				}
			}
		} else {
			ans = max(ans, right-left+1)
		}

		right++
	}
	return ans
}
