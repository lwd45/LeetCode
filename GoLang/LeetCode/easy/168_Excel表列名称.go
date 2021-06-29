package easy

func convertToTitle(columnNumber int) string {
	ans := []byte{}

	for columnNumber > 0 {
		temp := (columnNumber - 1) % 26
		ans = append(ans, 'A'+byte(temp))
		columnNumber = (columnNumber - temp - 1) / 26
	}

	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
