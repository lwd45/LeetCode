package mid

//
//func constructArr(a []int) []int {
//	ans, left, right := make([]int, len(a)), make([]int, len(a)), make([]int, len(a))
//	left[0], right[len(a)-1] = a[0], a[len(a)-1]
//	for i := 1; i < len(a); i++ {
//		left[i] = left[i-1] * a[i]
//		right[len(a)-1-i] = right[len(a)-i] * a[len(a)-1-i]
//	}
//
//	for i := 0; i < len(a); i++ {
//		num1, num2 := 1, 1
//		if i-1 >= 0 {
//			num1 = left[i-1]
//		}
//		if i+1 < len(a) {
//			num2 = right[i+1]
//		}
//		ans[i] = num1 * num2
//	}
//	return ans
//}

func constructArr(a []int) []int {
	if a == nil || len(a) == 0 {
		return a
	}

	ans := make([]int, len(a))
	ans[0] = 1
	for i := 1; i < len(a); i++ {
		ans[i] = ans[i-1] * a[i]
	}
	temp := 1
	for i := len(a) - 1; i >= 0; i-- {
		ans[i] *= temp
		temp *= a[i]
	}
	return ans
}
