package LeetCode

func countPairs(deliciousness []int) int {
	maxNum := deliciousness[0]
	for _, v := range deliciousness[1:] {
		if v > maxNum {
			maxNum = v
		}
	}

	maxSum := 2 * maxNum
	ans, cnt := 0, make(map[int]int)
	for _, v := range deliciousness {
		for i := 1; i <= maxSum; i = i << 1 {
			ans += cnt[i-v]
		}
		cnt[v]++
		ans = ans % 1000000007
	}
	return ans
}
