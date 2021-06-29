package mid

/*
给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
进阶：你可以在 O(n) 的时间解决这个问题吗？

输入：nums = [3,10,5,25,2,8]							//11 1010 101 11001 10 1000
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
输入：nums = [0]
输出：0
输入：nums = [2,4]									//10 100---110
输出：6
输入：nums = [8,10,2]								//1000 1010 10
输出：10
输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]	//
输出：127

1 <= nums.length <= 2 * 104
0 <= nums[i] <= 231 - 1
*/

func findMaximumXOR(nums []int) int {
	ans := 0
	for i := 30; i >= 0; i-- {
		set := make(map[int]struct{}, len(nums))
		for _, num := range nums {
			set[(num >> i)] = struct{}{}
		}

		ansNext, find := 2*ans+1, false
		for _, num := range nums {
			if _, ok := set[ansNext^(num>>i)]; ok {
				find = true
				break
			}
		}

		if find {
			ans = ansNext
		} else {
			ans = ansNext - 1
		}
	}
	return ans
}
