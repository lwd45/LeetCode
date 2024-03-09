package mid

//func divisibilityArray(word string, m int) []int {
//	ans := make([]int, len(word))
//
//	pre := 0
//	for i := 0; i < len(word); i++ {
//		now, _ := strconv.ParseInt(word[i:i+1], 10, 64)
//		pre = pre*10 + int(now)
//		if pre%m == 0 {
//			pre = 0
//			ans[i] = 1
//		} else {
//			pre = pre % m
//		}
//	}
//
//	return ans
//}

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))

	num := 0
	for i, c := range word {
		num = (num*10 + int(c-'0')) % m
		if num == 0 {
			ans[i] = 1
		}
	}

	return ans
}
