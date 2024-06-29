package mid

func smallestString(s string) string {
	var ans []byte

	flag := false
	for i := 0; i < len(s); i++ {
		if flag && s[i] == 'a' {
			ans = append(ans, s[i:]...)
			break
		} else if !flag && s[i] == 'a' {
			ans = append(ans, s[i])
			continue
		} else {
			flag = true
			ans = append(ans, s[i]-1)
		}
	}

	if !flag {
		ans = ans[:len(ans)-1]
		ans = append(ans, 'z')
	}
	return string(ans)
}
