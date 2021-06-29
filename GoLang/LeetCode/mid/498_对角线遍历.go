package mid

func findDiagonalOrder(mat [][]int) []int {
	row, col, ans := len(mat), len(mat[0]), make([]int, 0, len(mat)*len(mat[0]))

	i, j, isUp := 0, 0, true
	for len(ans) < cap(ans) {
		for 0 <= i && i < row && 0 <= j && j < col {
			ans = append(ans, mat[i][j])
			if isUp {
				i, j = i-1, j+1
			} else {
				i, j = i+1, j-1
			}
		}

		if isUp {
			if j >= col {
				j, i = col-1, i+2
			} else {
				i++
			}
		} else {
			if i >= row {
				i, j = row-1, j+2
			} else {
				j++
			}
		}
		isUp = !isUp
	}
	return ans
}
