package mid

func minOperations(nums []int) int {
	ans := 0

	opCnt := 0
	for _, num := range nums {
		if opCnt%2 == 0 { // 不变
			if num == 0 {
				ans++
				opCnt++
			}
		} else { // 相反
			if num == 1 {
				ans++
				opCnt++
			}
		}
	}

	return ans
}
