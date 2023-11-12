package mid

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	res := make([]int, len(potions), len(potions))
	for idx, spell := range spells {
		target := int(success-1)/spell + 1
		// 二分查找第一个 >= target 的数
		start, end := 0, len(potions)-1
		for start < end {
			mid := (start + end) / 2
			if potions[mid] >= target {
				end = mid
			} else {
				start = mid + 1
			}
		}

		if potions[start] >= target {
			res[idx] = len(potions) - start
		}
	}

	return res
}
