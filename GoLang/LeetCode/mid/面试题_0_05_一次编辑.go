package mid

/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
first = "pale"
second = "ple"
输出: True

first = "pales"
second = "pal"
输出: False
*/
//func oneEditAway(first string, second string) bool {
//	len1, len2 := len(first), len(second)
//	dp := make([][]int, len1+1)
//	for i := 0; i <= len1; i++ {
//		dp[i] = make([]int, len2+1)
//	}
//	for i := 0; i <= len1; i++ {
//		dp[i][0] = i
//	}
//	for i := 0; i <= len2; i++ {
//		dp[0][i] = i
//	}
//
//	for i := 1; i <= len1; i++ {
//		for j := 1; j <= len2; j++ {
//			if first[i-1] == second[j-1] {
//				dp[i][j] = dp[i-1][j-1]
//			} else {
//				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
//			}
//		}
//	}
//	return dp[len1][len2] <= 1
//}

func oneEditAway(first string, second string) bool {
	len1, len2 := len(first), len(second)
	if (len1-len2) > 1 || (len1-len2) < -1 {
		return false
	}
	if len1 < len2 {
		len1, len2, first, second = len2, len1, second, first
	}

	if len1 == len2 {
		count := 0
		for i := 0; i < len1; i++ {
			if first[i] != second[i] {
				count++
			}
		}
		return count < 2
	} else {
		for i := 0; i < len1; i++ {
			if i == 0 {
				if first[i+1:] == second[:] {
					return true
				}
			} else if i == len2 {
				if first[:i] == second[:] {
					return true
				}
			} else {
				if first[:i] == second[:i] && first[i+1:] == second[i:] {
					return true
				}
			}
		}
		return false
	}
}
