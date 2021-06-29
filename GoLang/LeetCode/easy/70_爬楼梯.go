package easy

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

输入： 2
输出： 2

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/
//func climbStairs(n int) int {
//	dp := make([]int, n+1)
//	dp[0], dp[1] = 1, 1
//	for i := 2; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[n]
//}

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	num1, num2 := 1, 1
	for i := 2; i <= n; i++ {
		num1, num2 = num2, num1+num2
	}
	return num2
}
