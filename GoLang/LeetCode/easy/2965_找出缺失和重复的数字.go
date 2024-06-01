package easy

func findMissingAndRepeatedValues(grid [][]int) []int {
	a, b := 0, 0
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row := (grid[i][j] - 1) / n
			col := (grid[i][j] - 1) - n*row

			if row == i && col == j {
				continue
			}

			if grid[row][col] == grid[i][j] {
				a = grid[i][j]
				b = i*n + j + 1
			} else {
				if a != 0 && grid[row][col] == a {
					b = i*n + j + 1
				}

				t := grid[i][j]
				grid[i][j] = grid[row][col]
				grid[row][col] = t
				j--
			}
		}
	}
	return []int{a, b}
}
