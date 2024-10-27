package mid

func maxScoreSightseeingPair(values []int) int {
	pos := 0
	posVal := values[0]

	ans := 0
	for i := 1; i < len(values); i++ {
		ans = max(ans, values[i]+posVal-(i-pos))

		if values[i] > values[pos]-(i-pos) {
			pos = i
			posVal = values[i]
		}
	}
	return ans
}
