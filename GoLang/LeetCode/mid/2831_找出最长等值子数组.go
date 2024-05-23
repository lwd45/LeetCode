package mid

//func longestEqualSubarray(nums []int, k int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	idxMap := make(map[int][]int)
//	for i, num := range nums {
//		idxMap[num] = append(idxMap[num], i)
//	}
//
//	ans := 1
//	for _, idxSlice := range idxMap {
//		left, right, lens := 0, 1, len(idxSlice)
//
//		gap := 0
//		for right < lens {
//			gap += idxSlice[right] - idxSlice[right-1] - 1
//			if gap > k {
//				for gap > k {
//					gap -= idxSlice[left+1] - idxSlice[left] - 1
//					left++
//				}
//			} else {
//				ans = max(ans, right-left+1)
//			}
//			right++
//		}
//	}
//	return ans
//}

//func longestEqualSubarray(nums []int, k int) int {
//	idxMap := make(map[int][]int)
//	for idx, num := range nums {
//		idxMap[num] = append(idxMap[num], idx)
//	}
//
//	ans := 0
//	for _, idxSlice := range idxMap {
//		left := 0
//		for right := 0; right < len(idxSlice); right++ {
//			for idxSlice[right]-idxSlice[left]-(right-left) > k {
//				left++
//			}
//			ans = max(ans, right-left+1)
//		}
//	}
//	return ans
//}

func longestEqualSubarray(nums []int, k int) int {
	cnt := make(map[int]int)

	l := 0
	ans := 0
	for r, num := range nums {
		cnt[num]++
		for (r-l+1)-cnt[nums[l]] > k {
			cnt[nums[l]]--
			l++
		}
		ans = max(ans, cnt[nums[r]])
	}
	return ans
}
