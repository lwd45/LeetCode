package mid

//func numSubarraysWithSum(nums []int, goal int) int {
//	cnt := make(map[int]int)
//	sum, ans := 0, 0
//	for _, v := range nums {
//		cnt[sum]++
//		sum += v
//		ans += cnt[sum-goal]
//	}
//	return ans
//}

func numSubarraysWithSum(nums []int, goal int) int {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0

	ans := 0
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}

		sum2 += num
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return ans
}
