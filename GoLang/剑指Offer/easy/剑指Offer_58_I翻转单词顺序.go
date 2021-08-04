package easy

import "strings"

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	ans := ""
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] == ""{
			continue
		}
		ans = ans + split[i] + " "
	}
	return strings.TrimSpace(s)
}
