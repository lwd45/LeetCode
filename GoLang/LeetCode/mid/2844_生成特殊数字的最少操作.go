package mid

func minimumOperations(num string) int {
	n := len(num)
	hasZero, hasFive := false, false

	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			if hasZero {
				return n - i - 2
			} else {
				hasZero = true
			}
		} else if num[i] == '2' {
			if hasFive {
				return n - i - 2
			}
		} else if num[i] == '5' {
			if hasZero {
				return n - i - 2
			}
			hasFive = true
		} else if num[i] == '7' {
			if hasFive {
				return n - i - 2
			}
		}
	}

	if hasZero {
		return n - 1
	}
	return n
}

// 00 25 50 75
