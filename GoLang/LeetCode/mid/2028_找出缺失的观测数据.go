package mid

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (m + n)

	sum := 0
	for _, roll := range rolls {
		sum += roll
	}

	if total < sum+1*n || total > sum+6*n {
		return nil
	}

	var ans []int

	left := total - sum
	for n > 0 {
		ans = append(ans, left/n)
		left -= left / n
		n--
	}
	return ans
}
