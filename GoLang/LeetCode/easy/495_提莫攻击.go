package easy

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 1 {
		return duration
	}
	result := 0
	prev := timeSeries[0]
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-prev > duration {
			result += duration
		} else {
			result += timeSeries[i] - prev
		}
		prev = timeSeries[i]
	}
	return result + duration
}
