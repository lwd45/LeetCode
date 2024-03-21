package mid

func maxMoves(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	res := make([][]int, row)
	for r := 0; r < row; r++ {
		res[r] = make([]int, col, col)
	}

	for r := 0; r < row; r++ {
		if r-1 >= 0 && grid[r][1] > grid[r-1][0] {
			res[r][1] = max(res[r][1], res[r-1][0]+1)
		}
		if grid[r][1] > grid[r][0] {
			res[r][1] = max(res[r][1], res[r][0]+1)
		}
		if r+1 < row && grid[r][1] > grid[r+1][0] {
			res[r][1] = max(res[r][1], res[r+1][0]+1)
		}
	}

	for c := 2; c < col; c++ {
		for r := 0; r < row; r++ {
			if r-1 >= 0 && grid[r][c] > grid[r-1][c-1] && res[r-1][c-1] != 0 {
				res[r][c] = max(res[r][c], res[r-1][c-1]+1)
			}
			if grid[r][c] > grid[r][c-1] && res[r][c-1] != 0 {
				res[r][c] = max(res[r][c], res[r][c-1]+1)
			}
			if r+1 < row && grid[r][c] > grid[r+1][c-1] && res[r+1][c-1] != 0 {
				res[r][c] = max(res[r][c], res[r+1][c-1]+1)
			}
		}
	}

	ans := 0
	for c := 1; c < col; c++ {
		for r := 0; r < row; r++ {
			ans = max(ans, res[r][c])
		}
	}
	return ans
}
