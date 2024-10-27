package mid

import "math"

func maxStrength(nums []int) int64 {
	negativeCnt, zeroCnt, positiveCnt := 0, 0, 0
	maxNegative := math.MinInt32

	ans := 1
	for _, num := range nums {
		if num > 0 {
			positiveCnt++
			ans *= num
		} else if num == 0 {
			zeroCnt++
		} else {
			negativeCnt++
			ans *= num
			if maxNegative < num {
				maxNegative = num
			}
		}
	}

	if negativeCnt == 1 && positiveCnt == 0 && zeroCnt == 0 {
		return int64(ans)
	} else if positiveCnt == 0 && negativeCnt <= 1 {
		return 0
	} else if ans > 0 {
		return int64(ans)
	} else {
		return int64(ans / maxNegative)
	}
}
