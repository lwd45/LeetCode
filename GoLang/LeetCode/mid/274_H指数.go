package mid

func hIndex(citations []int) int {
	n := len(citations)

	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation > n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}

	ans := 0
	for i := n; i > 0; i-- {
		ans += counter[i]
		if ans >= i {
			return i
		}
	}
	return 0
}

//func hIndex(citations []int) int {
//	sort.Slice(citations, func(i, j int) bool {
//		return citations[i] < citations[j]
//	})
//
//	idx := len(citations) - 1
//	for idx >= 0 && citations[idx] >= (len(citations)-idx) {
//		idx--
//	}
//
//	if idx < 0 {
//		return len(citations)
//	}
//
//	return len(citations) - 1 - idx
//}
