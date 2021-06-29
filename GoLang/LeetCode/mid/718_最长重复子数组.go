package mid

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) { //让nums2长度短
		nums1, nums2 = nums2, nums1
	}
	len1, len2, totalLen := len(nums1), len(nums2), len(nums1)+len(nums2)

	ret := 0
	for i := 1; i <= totalLen-1; i++ {
		start1, start2, lens := 0, 0, 0

		if i <= len2 {
			lens = i
			start2 = len2 - i
		} else if i <= len1 {
			lens = len2
			start1 = i - len2
		} else {
			lens = len2 - i + len1
			start1 = len1 - lens
		}

		ans := findMax(start1, start2, lens, &nums1, &nums2)
		if ans > ret {
			ret = ans
		}
	}
	return ret
}

func findMax(start1 int, start2 int, lens int, nums1, nums2 *[]int) int {
	max, now := 0, 0
	for i := 0; i < lens; i++ {
		if (*nums1)[start1+i] == (*nums2)[start2+i] {
			now++
			if now > max {
				max = now
			}
		} else {
			now = 0
		}
	}
	return max
}
