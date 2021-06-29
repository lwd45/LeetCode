package mid

/*
给你一个字符串 s，找到 s 中最长的回文子串。
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成
*/
func longestPalindrome(s string) string {
	start, end := 0, 0

	for i := 0; i < len(s)-1; i++ {
		i1, i2 := maxLen(s, i, i)
		i3, i4 := maxLen(s, i, i+1)
		if i2-i1 >= end-start {
			start, end = i1, i2
		}
		if i4-i3 >= end-start {
			start, end = i3, i4
		}
	}
	return s[start : end+1]
}

func maxLen(s string, i1 int, i2 int) (int, int) {
	for i1 >= 0 && i2 < len(s) && s[i1] == s[i2] {
		i1, i2 = i1-1, i2+1
	}
	return i1 + 1, i2 - 1
}
