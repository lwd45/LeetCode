package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	idx := m + n - 1
	for ; m-1 >= 0 && n-1 >= 0; {
		if nums2[n-1] > nums1[m-1] {
			nums1[idx] = nums2[n-1]
			n--
			idx--
		} else {
			nums1[idx] = nums1[m-1]
			m--
			idx--
		}
	}
	for ; n-1 >= 0; n-- {
		nums1[idx] = nums2[n-1]
		idx--
	}
}
