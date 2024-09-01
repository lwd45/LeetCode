package mid

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 0
	for i := 0; i < len(points); {
		ans++

		j := i + 1
		for ; j < len(points) && points[j][0] <= points[i][0]+w; j++ {
		}

		i = j
	}
	return ans
}
