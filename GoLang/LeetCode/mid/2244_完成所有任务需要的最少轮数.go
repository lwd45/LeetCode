package mid

import "math"

func minimumRounds(tasks []int) int {
	cnt := make(map[int]int, len(tasks))
	for _, task := range tasks {
		cnt[task]++
	}

	ans := 0
	for _, val := range cnt {
		if val == 1 {
			return -1
		}

		ans += int(math.Ceil(float64(val) / 3.0))
	}
	return ans
}
