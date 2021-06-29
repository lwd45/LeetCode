package mid

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

输入：nums = [0,1,0,3,2,3]
输出：4

输入：nums = [7,7,7,7,7,7,7]
输出：1

进阶：
你可以设计时间复杂度为 O(n2) 的解决方案吗？
你能将算法的时间复杂度降低到O(n log(n)) 吗?
*/
func lengthOfLIS(nums []int) int {
	length, arr := 1, make([]int, len(nums)+1)
	arr[length] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > arr[length] {
			length++
			arr[length] = nums[i]
		} else {
			left, right := 1, length
			for left < right {
				mid := (left + right) >> 1
				if arr[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			arr[left] = nums[i]
		}
	}
	return length
}
