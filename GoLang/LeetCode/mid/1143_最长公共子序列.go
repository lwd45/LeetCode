package mid

/*
给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

1 <= text1.length, text2.length <= 1000
text1 和text2 仅由小写英文字符组成。
*/
//func LongestCommonSubsequence(text1 string, text2 string) int {
//	len1, len2 := len(text1), len(text2)
//	dp := make([][]int, len1+1)
//	for i := range dp {
//		dp[i] = make([]int, len2+1)
//	}
//
//	maxAns := 0
//	for i := 0; i < len1; i++ {
//		for j := 0; j < len2; j++ {
//			if text2[j] == text1[i] {
//				dp[i+1][j+1] = dp[i][j] + 1
//				maxAns = max(maxAns, dp[i+1][j+1])
//			} else {
//				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1], dp[i][j])
//			}
//		}
//	}
//	return maxAns
//}
