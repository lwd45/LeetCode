package mid

import "sort"

/*
给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为m x n 由非负整数组成。
矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。
请你找出matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。

输入：matrix = [[5,2],[1,6]], k = 1
输出：7
解释：坐标 (0,1) 的值是 5 XOR 2 = 7 ，为最大的值。

输入：matrix = [[5,2],[1,6]], k = 2
输出：5
解释：坐标 (0,0) 的值是 5 = 5 ，为第 2 大的值。

输入：matrix = [[5,2],[1,6]], k = 3
输出：4
解释：坐标 (1,0) 的值是 5 XOR 1 = 4 ，为第 3 大的值。

输入：matrix = [[5,2],[1,6]], k = 4
输出：0
解释：坐标 (1,1) 的值是 5 XOR 2 XOR 1 XOR 6 = 0 ，为第 4 大的值。

m == matrix.length
n == matrix[i].length
1 <= m, n <= 1000
0 <= matrix[i][j] <= 106
1 <= k <= m * n
*/
func kthLargestValue(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}

	dp[0][0] = matrix[0][0]
	temp := make([]int, 0, row*col)
	temp = append(temp, dp[0][0])
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] ^ matrix[i][0]
		temp = append(temp, dp[i][0])
	}
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] ^ matrix[0][i]
		temp = append(temp, dp[0][i])
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = dp[i-1][j-1] ^ dp[i-1][j] ^ dp[i][j-1] ^ matrix[i][j]
			temp = append(temp, dp[i][j])
		}
	}
	sort.Ints(temp)
	return temp[len(temp)-k]
}
