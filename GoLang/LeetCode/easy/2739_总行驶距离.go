package easy

func distanceTraveled(mainTank int, additionalTank int) int {
	ans := 0
	for mainTank > 0 {
		if mainTank >= 5 && additionalTank > 0 {
			mainTank++
			additionalTank--
		}
		ans += 10 * min(5, mainTank)
		mainTank -= min(5, mainTank)
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
