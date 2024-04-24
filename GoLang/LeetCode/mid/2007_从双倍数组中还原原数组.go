package mid

import "sort"

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)

	count := make(map[int]int)
	for _, num := range changed {
		count[num]++
	}

	var res []int
	for _, num := range changed {
		if count[num] == 0 {
			continue
		}
		count[num]--

		if count[num*2] == 0 {
			return []int{}
		}
		count[num*2]--
		res = append(res, num)
	}
	return res
}
