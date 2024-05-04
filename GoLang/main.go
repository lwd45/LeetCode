package main

import (
	"strings"
)

func main() {
	baseNeg2(2)
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	builder := strings.Builder{}
	for n != 0 {
		r := n % -2
		if r == 0 {
			builder.WriteString("0")
		} else {
			n -= 1
			builder.WriteString("1")
		}
		n /= -2
	}
	return builder.String()
}
