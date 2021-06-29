package easy

/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。

0 ≤ x, y < 231.
输入: x = 1, y = 4

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置。
*/
func hammingDistance(x int, y int) int {
	temp := x ^ y
	ans := 0
	for temp > 0 {
		if temp&1 == 1 {
			ans++
		}
		temp = temp >> 1
	}
	return ans
}
