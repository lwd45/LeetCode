package mid

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans := 0
	leftA, leftB := capacityA, capacityB
	left, right := 0, len(plants)-1

	for left < right {
		if plants[left] > leftA {
			leftA = capacityA
			ans++
		}
		leftA -= plants[left]

		if plants[right] > leftB {
			leftB = capacityB
			ans++
		}
		leftB -= plants[right]

		left++
		right--
	}

	if left == right && max(leftA, leftB) < plants[left] {
		ans++
	}
	return ans
}
