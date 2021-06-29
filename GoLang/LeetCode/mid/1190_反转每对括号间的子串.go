package mid

func reverseParentheses(s string) string {
	stack, stackIndex := make([]uint8, len(s)), 0

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			var str []uint8
			for stack[stackIndex-1] != '(' {
				str = append(str, stack[stackIndex-1])
				stackIndex--
			}
			stackIndex--
			for j := 0; j < len(str); j++ {
				stack[stackIndex] = str[j]
				stackIndex++
			}
		} else {
			stack[stackIndex] = s[i]
			stackIndex++
		}
	}

	stack = stack[:stackIndex]
	return string(stack)
}
