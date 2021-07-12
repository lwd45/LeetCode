package mid

func HIndex(citations []int) int {
	left, right := 0, len(citations)-1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] >= len(citations)-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(citations) - left
}
