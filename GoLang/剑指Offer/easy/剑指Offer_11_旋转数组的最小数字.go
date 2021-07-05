package easy

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) >> 1
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right-- //因为相等，则左侧至少有一个替代品
		}
	}
	return numbers[left]
}
