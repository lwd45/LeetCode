package easy

//func isCovered(ranges [][]int, l int, r int) bool {
//	newRanges := make([][]int, 0)
//
//	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
//	left, right := ranges[0][0], ranges[0][1]
//	for i := 1; i < len(ranges); i++ {
//		if ranges[i][0] > right+1 {
//			newRanges = append(newRanges, []int{left, right})
//			left, right = ranges[i][0], ranges[i][1]
//		} else if ranges[i][1] > right {
//			right = ranges[i][1]
//		}
//	}
//	newRanges = append(newRanges, []int{left, right})
//
//	for _, newRange := range newRanges {
//		if l < newRange[0] {
//			return false
//		} else if l <= newRange[1] {
//			l = newRange[1] + 1
//		}
//
//		if l > r {
//			return true
//		}
//	}
//	return l > r
//}
func isCovered(ranges [][]int, l int, r int) bool {
	diff := [52]int{}
	for _, rg := range ranges {
		diff[rg[0]]++
		diff[rg[1]+1]--
	}

	curr := 0
	for i := 1; i <= 51; i++ {
		curr += diff[i]
		if l <= i && i <= r && curr < 0 {
			return false
		}
	}
	return true
}
