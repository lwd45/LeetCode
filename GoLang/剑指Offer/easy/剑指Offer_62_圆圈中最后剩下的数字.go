package easy

/*
0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

输入: n = 5, m = 3
输出:3
示例 2：

输入: n = 10, m = 17
输出:2
*/

//func lastRemaining(n int, m int) int {
//	nums := make([]int, n, n)
//	for i := 0; i < n; i++ {
//		nums[i] = i
//	}
//	startIndex := 0
//	for len(nums) > 1 {
//		removeIndex := ((m-1)%len(nums) + startIndex) % len(nums)
//		nums = append(nums[:removeIndex], nums[removeIndex+1:]...)
//		startIndex = removeIndex
//	}
//	return nums[0]
//}

func lastRemaining(n int, m int) int {
	start := 0
	for i := 2; i <= n; i++ {
		start = (start + m) % i
	}
	return start
}
