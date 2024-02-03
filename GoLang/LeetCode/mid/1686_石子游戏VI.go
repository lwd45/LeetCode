package mid

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	//a[i]-b[j]-a[j]+b[i]>0 时 a选i
	//
	//b[i]-a[j]-b[j]+a[i]>0 时 b选i

	n := len(aliceValues)
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		values[i] = []int{aliceValues[i] + bobValues[i], aliceValues[i], bobValues[i]}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i][0] > values[j][0]
	})

	a, b := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a += values[i][1]
		} else {
			b += values[i][2]
		}
	}

	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
