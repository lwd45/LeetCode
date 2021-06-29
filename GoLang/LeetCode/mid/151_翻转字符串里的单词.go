package mid

import "strings"

/*
"the sky is blue"
"eht yks si eulb"

func reverseWords(s string) string {
	res := strings.Fields(s)
	length := len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return strings.Join(res, " ")
}
*/
func reverseWords(s string) string {
	fields := strings.Fields(s)

	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	return strings.Join(fields, " ")
}
