package mid

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	var ans float64 = 1
	for n != 0 {
		if (n & 1) == 1 {
			ans *= x
		}
		x *= x
		n = n >> 1
	}
	return ans
}
