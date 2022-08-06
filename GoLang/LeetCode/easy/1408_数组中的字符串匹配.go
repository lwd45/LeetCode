package easy

import "strings"

func stringMatching(words []string) []string {
	var res []string

	for i, wordI := range words {
		for j, wordJ := range words {
			if i != j && strings.Contains(wordJ, wordI) {
				res = append(res, wordI)
				break
			}
		}
	}
	return res
}
