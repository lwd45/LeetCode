package mid

//func maximumBeauty(nums []int, k int) int {
//	sort.Ints(nums)
//
//	ans := 0
//	i, j := 0, 0
//	for ; j < len(nums); j++ {
//		for nums[j]-nums[i] > 2*k {
//			i++
//		}
//		ans = max(ans, j-i+1)
//	}
//	return ans
//}

func maximumBeauty(nums []int, k int) int {
	cnt := make([]int, 100002)

	for _, num := range nums {
		l, r := max(num-k, 0), min(num+k, 100000)
		cnt[l]++
		cnt[r+1]--
	}

	ans := 0
	sum := 0
	for _, n := range cnt {
		sum += n
		ans = max(sum, ans)
	}
	return ans
}
