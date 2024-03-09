package mid

func minimumPossibleSum(n int, target int) int {
	mod := 1000000007

	m := target / 2
	if n <= m {
		return ((1 + n) * n / 2) % mod
	}

	sum1 := (1 + m) * m / 2
	sum2 := (target + target + n - m - 1) * (n - m) / 2
	return (sum1 + sum2) % mod
}
