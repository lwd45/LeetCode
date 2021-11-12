package mid

import "math"

/*
我们正在玩一个猜数游戏，游戏规则如下：
我从1到 n 之间选择一个数字。
你来猜我选了哪个数字。
如果你猜到正确的数字，就会 赢得游戏 。
如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。
给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。
*/
//var dp [201][201]int
//
//func getMoneyAmount(n int) int {
//	return dfs(1, n)
//}
//
//func dfs(left int, right int) int {
//	if left >= right {
//		return 0
//	}
//	if dp[left][right] != 0 {
//		return dp[left][right]
//	}
//
//	maxAns := math.MaxInt64
//	for i := left; i <= right; i++ {
//		cur := max(dfs(left, i-1), dfs(i+1, right)) + i
//		maxAns = min(maxAns, cur)
//	}
//	dp[left][right] = maxAns
//	return maxAns
//}

/*
public int getMoneyAmount(int n) {
	int[][] f = new int[n + 10][n + 10];
	for (int len = 2; len <= n; len++) {
		for (int l = 1; l + len - 1 <= n; l++) {
			int r = l + len - 1;
			f[l][r] = 0x3f3f3f3f;
			for (int x = l; x <= r; x++) {
				int cur = Math.max(f[l][x - 1], f[x + 1][r]) + x;
				f[l][r] = Math.min(f[l][r], cur);
			}
		}
	}
	return f[1][n];
}
*/
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for lens := 2; lens <= n; lens++ {
		for l := 1; l+lens-1 <= n; l++ {
			r := l + lens - 1
			dp[l][r] = math.MaxInt64
			for i := l; i <= r; i++ {
				cur := max(dp[l][i-1], dp[i+1][r]) + i
				dp[l][r] = min(dp[l][r], cur)
			}
		}
	}
	return dp[1][n]
}
