package mid

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
//func maxProduct(nums []int) int {
//	if nums == nil || len(nums) == 0 {
//		return 0
//	}
//	sMax, sMin := make([]int, len(nums)), make([]int, len(nums))
//	sMax[0], sMin[0] = nums[0], nums[0]
//
//	ret := nums[0]
//	for i := 1; i < len(nums); i++ {
//		num1 := sMax[i-1] * nums[i]
//		num2 := sMin[i-1] * nums[i]
//		sMax[i] = Max152(num1, num2, nums[i])
//		sMin[i] = Min152(num1, num2, nums[i])
//		ret = Max152(sMax[i], sMin[i], ret)
//	}
//	return ret
//}
//
//func Max152(nums ...int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}
//
//func Min152(nums ...int) int {
//	min := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] < min {
//			min = nums[i]
//		}
//	}
//	return min
//}
