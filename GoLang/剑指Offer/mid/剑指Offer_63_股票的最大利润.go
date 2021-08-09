package mid

//func maxProfit(prices []int) int {
//	ans := 0
//	slice := make([]int, 0)
//	for _, price := range prices {
//		if len(slice) > 0 && price <= slice[len(slice)-1] {
//			for len(slice) > 0 && price <= slice[len(slice)-1] {
//				slice = slice[:len(slice)-1]
//			}
//		} else if len(slice) > 0 {
//			ans = max_63(ans, price-slice[0])
//		} else {
//			slice = append(slice, price)
//		}
//	}
//	return ans
//}
//
func max_63(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min_63(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	ans, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max_63(ans, prices[i]-min)
		min = min_63(min, prices[i])
	}
	return ans
}
