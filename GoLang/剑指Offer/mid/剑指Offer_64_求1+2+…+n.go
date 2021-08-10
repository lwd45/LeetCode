package mid

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(i int) bool {
		ans += i
		return i > 0 && sum(i-1)
	}
	sum(n)
	return ans
}
