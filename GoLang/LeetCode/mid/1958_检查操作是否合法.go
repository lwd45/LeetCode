package mid

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	//上
	r, c := rMove-1, cMove
	otherColor := 0
	for ; r >= 0; r-- {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//下
	r, c = rMove+1, cMove
	otherColor = 0
	for ; r < len(board); r++ {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//左
	r, c = rMove, cMove-1
	otherColor = 0
	for ; c >= 0; c-- {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//右
	r, c = rMove, cMove+1
	otherColor = 0
	for ; c < len(board[0]); c++ {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//左上
	r, c = rMove-1, cMove-1
	otherColor = 0
	for ; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//左下
	r, c = rMove+1, cMove-1
	otherColor = 0
	for ; r < len(board) && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//右下
	r, c = rMove+1, cMove+1
	otherColor = 0
	for ; r < len(board) && c < len(board[0]); r, c = r+1, c+1 {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	//右上
	r, c = rMove-1, cMove+1
	otherColor = 0
	for ; r >= 0 && c < len(board[0]); r, c = r-1, c+1 {
		if board[r][c] == '.' || (otherColor == 0 && board[r][c] == color) {
			break
		} else if board[r][c] != color {
			otherColor++
			continue
		}
		return true
	}

	return false
}
