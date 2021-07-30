package easy

func titleToNumber(columnTitle string) int {
	sum := 0
	base := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		sum += (int(columnTitle[i]-'A') + 1) * base
		base *= 26
	}
	return sum
}
