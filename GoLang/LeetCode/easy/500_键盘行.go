package easy

import "strings"

func findWords(words []string) []string {
	ans := []string{}
	for _, word := range words {
		preIndex := 0
		tempWord := strings.ToLower(word)
		for i, c := range tempWord {
			switch c {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				if preIndex == 0 || preIndex == 1 {
					preIndex = 1
					if i == len(word)-1 {
						ans = append(ans, word)
					}
				} else {
					goto End
				}
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				if preIndex == 0 || preIndex == 2 {
					preIndex = 2
					if i == len(word)-1 {
						ans = append(ans, word)
					}
				} else {
					goto End
				}
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				if preIndex == 0 || preIndex == 3 {
					preIndex = 3
					if i == len(word)-1 {
						ans = append(ans, word)
					}
				} else {
					goto End
				}
			}
		}
	End:
	}
	return ans
}
