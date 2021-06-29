package mid

func snakesAndLadders(board [][]int) int {
	lens := len(board)
	nums := make([]int, 0, lens*lens+1)
	nums = append(nums, -1)

	left2Right := true
	for i := lens - 1; i >= 0; i-- {
		if left2Right {
			for j := 0; j < lens; j++ {
				nums = append(nums, board[i][j])
			}
		} else {
			for j := lens - 1; j >= 0; j-- {
				nums = append(nums, board[i][j])
			}
		}
		left2Right = !left2Right
	}

	queue := make([]int, 0)
	queue = append(queue, 1)
	maps := map[int]int{}
	maps[1] = 0

	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]

		for i := 1; i <= 6; i++ {
			next := now + i
			if next > lens*lens {
				break
			}
			if nums[next] != -1 {
				next = nums[next]
			}
			if next == lens*lens {
				return maps[now] + 1
			}
			if _, ok := maps[next]; !ok {
				maps[next] = maps[now] + 1
				queue = append(queue, next)
			}
		}
	}
	return -1
}
