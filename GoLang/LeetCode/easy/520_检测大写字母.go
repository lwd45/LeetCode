package easy

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	if 'A' <= word[0] && word[0] <= 'Z' { //首字母大写
		if 'A' <= word[1] && word[1] <= 'Z' {
			for i := 2; i < len(word); i++ {
				if 'a' <= word[i] && word[i] <= 'z' {
					return false
				}
			}
			return true
		} else {
			for i := 2; i < len(word); i++ {
				if 'A' <= word[i] && word[i] <= 'Z' {
					return false
				}
			}
			return true
		}
	} else { //首字母小写
		for _, v := range word {
			if 'A' <= v && v <= 'Z' {
				return false
			}
		}
		return true
	}
}
