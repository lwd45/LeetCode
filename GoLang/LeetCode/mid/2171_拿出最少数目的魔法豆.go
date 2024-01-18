package mid

import "sort"

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)

	total := 0
	for _, v := range beans {
		total += v
	}

	res := total
	for i := 0; i < n; i++ {
		res = min(res, total-beans[i]*(n-i))
	}
	return int64(res)
}
