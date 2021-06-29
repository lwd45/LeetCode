package mid

func findMaxLength(nums []int) int {
	max := 0
	maps := make(map[int]int, len(nums)) //保存sum的index
	maps[0] = -1

	sum := 0
	for index, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}

		preIndex, ok := maps[sum]
		if ok {
			if (index - preIndex) > max {
				max = index - preIndex
			}
		} else {
			maps[sum] = index
		}
	}
	return max
}
