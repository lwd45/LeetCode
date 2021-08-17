package easy

func checkRecord(s string) bool {
	absentCount, lateCount := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			lateCount++
			if lateCount >= 3 {
				return false
			}
		} else if s[i] == 'A' {
			absentCount++
			if absentCount >= 2 {
				return false
			}
			lateCount = 0
		} else {
			lateCount = 0
		}
	}
	return true
}
