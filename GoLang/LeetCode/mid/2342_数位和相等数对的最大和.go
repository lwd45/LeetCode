package mid

//func maximumSum(nums []int) int {
//	m := make(map[int][2]int)
//
//	ans := -1
//	for _, num := range nums {
//		sum := 0
//
//		t := num
//		for t > 0 {
//			sum += t % 10
//			t = t / 10
//		}
//
//		if _, ok := m[sum]; !ok {
//			m[sum] = [2]int{}
//		}
//
//		if num > m[sum][0] {
//			m[sum] = [2]int{num, m[sum][0]}
//			if m[sum][1] != 0 && m[sum][0]+m[sum][1] > ans {
//				ans = m[sum][0] + m[sum][1]
//			}
//		} else if num > m[sum][1] {
//			m[sum] = [2]int{m[sum][0], num}
//			if m[sum][0]+m[sum][1] > ans {
//				ans = m[sum][0] + m[sum][1]
//			}
//		}
//	}
//	return ans
//}
