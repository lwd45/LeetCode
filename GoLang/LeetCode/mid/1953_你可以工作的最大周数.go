package mid

func numberOfWeeks(milestones []int) int64 {
	if len(milestones) == 1 {
		return 1
	}

	sum := 0
	maxVal := milestones[0]
	for _, milestone := range milestones {
		sum += milestone
		maxVal = max(maxVal, milestone)
	}

	if sum-maxVal >= maxVal {
		return int64(sum)
	}
	return int64((sum-maxVal)*2 + 1)
}
