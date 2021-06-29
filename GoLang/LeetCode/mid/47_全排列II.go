package mid

import "sort"

//func PermuteUnique(nums []int) (ans [][]int) {
//	sort.Ints(nums)
//	vis := make([]bool, len(nums))
//	var now []int
//
//	var backTrace func(int)
//	backTrace = func(idx int) {
//		if idx == len(nums) {
//			ans = append(ans, append([]int{}, now...))
//			return
//		}
//
//		for i, v := range nums {
//			if vis[i] || (i > 0 && !vis[i-1] && v == nums[i-1]) {
//				continue
//			}
//			vis[i] = true
//			now = append(now, v)
//			backTrace(idx + 1)
//			vis[i] = false
//			now = now[:len(now)-1]
//		}
//	}
//	backTrace(0)
//	return
//}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	vis := make([]bool, len(nums))
	var now []int

	var backTrace func(idx int)
	backTrace = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, now...))
			return
		}
		for i, v := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && v == nums[i-1]) {
				continue
			}

			vis[i] = true
			now = append(now, v)
			backTrace(idx + 1)
			vis[i] = false
			now = now[:len(now)-1]
		}
	}
	backTrace(0)
	return
}
