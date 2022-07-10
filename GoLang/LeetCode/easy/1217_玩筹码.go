package easy

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, num := range position {
		if num%2 != 0 {
			odd += 1
		} else {
			even += 1
		}
	}

	if odd > even {
		return even
	}
	return odd
}
