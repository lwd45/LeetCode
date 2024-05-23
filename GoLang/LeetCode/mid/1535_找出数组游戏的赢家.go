package mid

func getWinner(arr []int, k int) int {
	if k == 1 {
		return max(arr[0], arr[1])
	}

	big := max(arr[0], arr[1])
	time := 1
	for i := 2; i < len(arr); i++ {
		if time >= k {
			return big
		}
		t := max(big, arr[i])
		if t == big {
			time++
		} else {
			big = t
			time = 1
		}
	}
	return big
}
