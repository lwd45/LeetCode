package mid

//mat := [][]int{{4, 3, 5}, {1, 2, 6}}
func FirstCompleteIndex(arr []int, mat [][]int) int {
	row, col := len(mat), len(mat[0])

	rowMap, colMap := make(map[int]int, len(arr)), make(map[int]int, len(arr))
	for r, nums := range mat {
		for c, num := range nums {
			rowMap[num] = r
			colMap[num] = c
		}
	}

	rMap, cMap := make(map[int]int, row), make(map[int]int, col)
	for idx, num := range arr {
		rMap[rowMap[num]]++
		cMap[colMap[num]]++
		if rMap[rowMap[num]] == col || cMap[colMap[num]] == row {
			return idx
		}
	}
	return -1
}
