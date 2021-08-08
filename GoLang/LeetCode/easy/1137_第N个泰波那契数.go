package easy

func tribonacci(n int) int {
	temp := [38]int{0, 1, 1}
	for i := 3; i <= n; i++ {
		temp[i] = temp[i-1] + temp[i-2] + temp[i-3]
	}
	return temp[n]
}
