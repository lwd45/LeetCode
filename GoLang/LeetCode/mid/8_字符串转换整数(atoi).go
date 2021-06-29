package mid

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	switch s[0] {
	case '+':
		return getNum(1, s[1:])
	case '-':
		return getNum(-1, s[1:])
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return getNum(1, s)
	default:
		return 0
	}
}

func getNum(sign int, s string) int {
	ans := 0
	for _, c := range s {
		if '0' <= c && c <= '9' {
			ans = ans*10 + int(c-'0')

			if sign == -1 && -ans <= math.MinInt32 {
				return math.MinInt32
			} else if sign == 1 && ans >= math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	return ans * sign
}
