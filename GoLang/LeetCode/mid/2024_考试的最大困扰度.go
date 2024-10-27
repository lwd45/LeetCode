package mid

func maxConsecutiveAnswers(answerKey string, k int) int {
	ans := 0

	t := k
	start, end := 0, 0
	for ; end < len(answerKey); end++ {
		if answerKey[end] == 'T' {
			ans = max(ans, end-start+1)
		} else if k > 0 {
			k--
			ans = max(ans, end-start+1)
		} else {
			for ; start <= end && answerKey[start] == 'T'; start++ {
			}
			start++
		}
	}

	k = t
	start, end = 0, 0
	for ; end < len(answerKey); end++ {
		if answerKey[end] == 'F' {
			ans = max(ans, end-start+1)
		} else if k > 0 {
			k--
			ans = max(ans, end-start+1)
		} else {
			for ; start <= end && answerKey[start] == 'F'; start++ {
			}
			start++
		}
	}

	return ans
}
