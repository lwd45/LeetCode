package mid

import "sort"

func diagonalSort(mat [][]int) [][]int {
	row, col := len(mat), len(mat[0])

	//	先处理列
	for i := 0; i < col; i++ {
		r, c := 0, i

		temp := make([]int, 0, row+col)
		for r < row && c < col {
			temp = append(temp, mat[r][c])
			r++
			c++
		}
		sort.Ints(temp)

		r, c = 0, i
		for r < row && c < col {
			mat[r][c] = temp[0]
			r++
			c++
			temp = temp[1:]
		}
	}

	//	再处理行
	for i := 1; i < row; i++ {
		r, c := i, 0

		temp := make([]int, 0, row+col)
		for r < row && c < col {
			temp = append(temp, mat[r][c])
			r++
			c++
		}
		sort.Ints(temp)

		r, c = i, 0
		for r < row && c < col {
			mat[r][c] = temp[0]
			r++
			c++
			temp = temp[1:]
		}
	}
	return mat
}
