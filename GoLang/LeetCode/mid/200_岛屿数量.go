package mid

func numIslands(grid [][]byte) int {
	row, col, ans := len(grid), len(grid[0]), 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs200(&grid, i, j)
			}
		}
	}
	return ans
}

func dfs200(grid *[][]byte, i int, j int) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	dfs200(grid, i+1, j)
	dfs200(grid, i-1, j)
	dfs200(grid, i, j+1)
	dfs200(grid, i, j-1)
}
