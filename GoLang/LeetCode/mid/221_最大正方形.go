package mid

func maximalSquare(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	ret := 0
	for i := 0; i < row; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ret = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			ret = 1
		} else {
			dp[0][i] = 0
		}
	}

	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if matrix[r][c] == '1' {
				dp[r][c] = min(dp[r-1][c-1], dp[r-1][c], dp[r][c-1]) + 1
				ret = max(ret, dp[r][c])
			} else {
				dp[r][c] = 0
			}
		}
	}
	return ret * ret
}
