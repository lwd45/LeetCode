package easy

//[0,0,1,-1]
func captureForts(forts []int) int {
	ans := 0
	for start, end := 0, 1; end < len(forts); end++ {
		if forts[end] != 0 {
			if forts[start] != 0 && forts[start] != forts[end] {
				if end-start-1 > ans {
					ans = end - start - 1
				}
			}
			start = end
		}
	}
	return ans
}

//func captureForts(forts []int) int {
//	start, ans := -1, 0
//	for idx, val := range forts {
//		if val != 0 {
//			if start != -1 && forts[start] != val && idx-start-1 > ans {
//				ans = idx - start - 1
//			}
//		}
//		start = idx
//	}
//	return ans
//}
