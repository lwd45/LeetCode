package easy

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}
	minRow, maxRow, minCol, maxCol := 0, len(matrix)-1, 0, len(matrix[0])-1
	direc, dIdx := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}, 0
	curRow, curCol := 0, 0

	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	for minRow <= maxRow && minCol <= maxCol {
		if curRow >= minRow && curRow <= maxRow && curCol >= minCol && curCol <= maxCol {
			ans = append(ans, matrix[curRow][curCol])
			curRow += direc[dIdx][0]
			curCol += direc[dIdx][1]
		} else {
			curRow -= direc[dIdx][0]
			curCol -= direc[dIdx][1]
			switch dIdx {
			case 0:
				minRow++
			case 1:
				maxCol--
			case 2:
				maxRow--
			case 3:
				minCol++
			}
			dIdx = (dIdx + 1) % 4
			curRow += direc[dIdx][0]
			curCol += direc[dIdx][1]
		}
	}
	return ans
}
