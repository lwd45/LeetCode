package hard

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	len := len1 + len2

	var (
		mid1   int
		mid2   int
		value1 int
		value2 int
	)
	if len%2 == 0 {
		mid1 = len/2 - 1
		mid2 = len / 2
	} else {
		mid1 = len / 2
		mid2 = len / 2
	}

	index1 := 0
	index2 := 0
	for i := 0; i < len; i++ {
		if i > mid2 {
			break
		}

		if i == mid1 && i == mid2 {
			if index1 < len1 && index2 < len2 {
				if nums1[index1] <= nums2[index2] {
					value1 = nums1[index1]
					index1++
				} else {
					value1 = nums2[index2]
					index2++
				}
			} else if index1 < len1 {
				value1 = nums1[index1]
				index1++
			} else {
				value1 = nums2[index2]
				index2++
			}
			value2 = value1
			break
		} else if i == mid1 {
			if index1 < len1 && index2 < len2 {
				if nums1[index1] <= nums2[index2] {
					value1 = nums1[index1]
					index1++
				} else {
					value1 = nums2[index2]
					index2++
				}
			} else if index1 < len1 {
				value1 = nums1[index1]
				index1++
			} else {
				value1 = nums2[index2]
				index2++
			}
		} else if i == mid2 {
			if index1 < len1 && index2 < len2 {
				if nums1[index1] <= nums2[index2] {
					value2 = nums1[index1]
					index1++
				} else {
					value2 = nums2[index2]
					index2++
				}
			} else if index1 < len1 {
				value2 = nums1[index1]
				index1++
			} else {
				value2 = nums2[index2]
				index2++
			}
		} else {
			if index1 < len1 && index2 < len2 {
				if nums1[index1] <= nums2[index2] {
					index1++
				} else {
					index2++
				}
			} else if index1 < len1 {
				index1++
			} else {
				index2++
			}
		}
	}
	return float64(value1+value2) / 2
}
