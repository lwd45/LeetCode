package mid

func countBattleships(board [][]byte) int {
	ans := 0

	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'X' &&
				(i == 0 || board[i-1][j] != 'X') &&
				(j == 0 || board[i][j-1] != 'X') {
				ans++
			}
		}
	}

	return ans
}
