package mid

/*
4		0100
14		1110
2		0010
4 		0100


*/
func totalHammingDistance(nums []int) int {
	var ans = 0
	for i := 0; i < 31; i++ {
		c := 0
		for _, num := range nums {
			c += (num >> i) & 1
		}
		ans += (len(nums) - c) * c
	}
	return ans
}
