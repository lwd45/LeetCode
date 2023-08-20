package easy

func circularGameLosers(n int, k int) []int {
	visit := make([]bool, n)
	idx := 0
	for i := k; !visit[idx]; i += k {
		visit[idx] = true
		idx = (idx + i) % n
	}

	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}

//func circularGameLosers(n int, k int) []int {
//	cntMap := make(map[int]int)
//	cntMap[1] = 1
//	idx, time := 1, 1
//
//	for {
//		idx = idx + time*k
//		if idx%n == 0 {
//			idx = n
//		} else {
//			idx = idx % n
//		}
//
//		if _, ok := cntMap[idx]; ok {
//			break
//		}
//		cntMap[idx]++
//		time++
//	}
//
//	ans := make([]int, 0)
//	for i := 1; i <= n; i++ {
//		if _, ok := cntMap[i]; !ok {
//			ans = append(ans, i)
//		}
//	}
//	return ans
//}
