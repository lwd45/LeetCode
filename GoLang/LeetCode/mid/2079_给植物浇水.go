package mid

func wateringPlants(plants []int, capacity int) int {
	ans := 0

	nowIdx := -1
	left := capacity
	for idx, val := range plants {
		if left < val {
			ans += nowIdx + 1
			nowIdx = -1
			left = capacity
		}
		ans += idx - nowIdx
		nowIdx = idx
		left -= val
	}
	return ans
}
