package hard

func reversePairs(nums []int) int {
	temp := make([]int, len(nums))
	ans := 0
	mergeSort(&nums, &temp, &ans, 0, len(nums)-1)
	return ans
}

func mergeSort(nums *[]int, temp *[]int, ans *int, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	mergeSort(nums, temp, ans, left, mid)
	mergeSort(nums, temp, ans, mid+1, right)

	for i := left; i <= right; i++ {
		(*temp)[i] = (*nums)[i]
	}
	l, r := left, mid+1
	for i := left; i <= right; i++ {
		if l == mid+1 {
			(*nums)[i] = (*temp)[r]
			r++
		} else if r == right+1 || (*temp)[l] <= (*temp)[r] {
			(*nums)[i] = (*temp)[l]
			l++
		} else {
			(*nums)[i] = (*temp)[r]
			r++
			*ans += mid + 1 - l
		}
	}
}
