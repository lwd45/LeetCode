package easy

func maxCount(m int, n int, ops [][]int) int {
	minC, minR := m, n
	for _, op := range ops {
		if op[0] < minC {
			minC = op[0]
		}
		if op[1] < minR {
			minR = op[1]
		}
	}
	return minC * minR
}
