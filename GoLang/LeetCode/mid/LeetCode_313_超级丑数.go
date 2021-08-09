package mid

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp, pointers := make([]int, n), make([]int, len(primes))
	dp[0] = 1

	for i := 1; i < n; i++ {
		minVal := math.MaxInt64
		for j := 0; j < len(pointers); j++ {
			minVal = min_313(minVal, dp[pointers[j]]*primes[j])
		}
		for j := 0; j < len(pointers); j++ {
			if minVal/dp[pointers[j]] == primes[j] {
				pointers[j] ++
			}
		}
		dp[i] = minVal
	}
	return dp[n-1]
}

func min_313(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
