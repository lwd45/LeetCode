package mid

func groupThePeople(groupSizes []int) [][]int {
	sizeMap := make(map[int][]int)
	for i, size := range groupSizes {
		sizeMap[size] = append(sizeMap[size], i)
	}

	var res [][]int
	for size, idxList := range sizeMap {
		start, end := 0, size
		for end <= len(idxList) {
			res = append(res, idxList[start:end])
			start += size
			end += size
		}
	}
	return res
}
