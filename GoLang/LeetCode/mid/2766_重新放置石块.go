package mid

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	stoneMap := make(map[int]int)
	for _, num := range nums {
		stoneMap[num] = 1
	}
	for i := 0; i < len(moveFrom); i++ {
		from, to := moveFrom[i], moveTo[i]
		stoneMap[from] = 0
		stoneMap[to] = 1
	}

	ans := make([]int, 0, len(stoneMap))
	for idx, cnt := range stoneMap {
		if cnt != 0 {
			ans = append(ans, idx)
		}
	}
	sort.Ints(ans)
	return ans
}
