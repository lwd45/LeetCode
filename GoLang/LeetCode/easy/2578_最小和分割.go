package easy

import (
	"sort"
	"strconv"
)

//func splitNum(num int) int {
//	var nums []int
//	for num > 0 {
//		i := num % 10
//		num = num / 10
//		nums = append(nums, i)
//	}
//
//	num1, num2 := 0, 0
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//
//	for _, v := range nums {
//		if num1 < num2 {
//			num1 = num1*10 + v
//		} else {
//			num2 = num2*10 + v
//		}
//	}
//	return num1 + num2
//}

func splitNum(num int) int {
	bytes := []byte(strconv.Itoa(num))
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })

	ans := [2]int{}
	for idx, v := range bytes {
		ans[idx&1] = ans[idx&1]*10 + int(v-'0')
	}
	return ans[0] + ans[1]
}
