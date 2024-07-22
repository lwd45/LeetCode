package mid

import "math"

func minimumMoves(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	var more, less []points
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] > 1 {
				for k := 0; k < grid[i][j]-1; k++ {
					more = append(more, points{x: i, y: j})
				}
			} else if grid[i][j] == 0 {
				less = append(less, points{x: i, y: j})
			}
		}
	}

	ans := math.MaxFloat64
	for i := 0; i < len(more); i++ {
		for j := i + 1; j < len(more); j++ {
			more[i], more[j] = more[j], more[i]
			ans = math.Min(ans, distance(more, less))
			more[i], more[j] = more[j], more[i]
		}
	}
	return int(ans)
}

func distance(more, less []points) float64 {
	dis := 0.0
	for i := 0; i < len(more); i++ {
		dis += math.Abs(float64(more[i].x-less[i].x)) + math.Abs(float64(more[i].y-less[i].y))
	}
	return dis
}

type points struct {
	x int
	y int
}
