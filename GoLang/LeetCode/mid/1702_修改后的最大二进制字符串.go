package mid

func maximumBinaryString(binary string) string {
	n := len(binary)
	s := []rune(binary)

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			for j := i + 1; j < n; j++ {
				if s[j] == '0' {
					s[j] = '1'
					s[i] = '1'
					s[i+1] = '0'
					i = i + 1
				}
			}
			break
		}
	}
	return string(s)
}
