package easy

import "strconv"

func summaryRanges(nums []int) []string {
	var res []string
	for i := 0; i < len(nums); {
		left := i
		for i++; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		}
		if left < i-1 {
			str := strconv.Itoa(nums[left]) + "->" + strconv.Itoa(nums[i-1])
			res = append(res, str)
		} else {
			res = append(res, strconv.Itoa(nums[left]))
		}
	}
	return res
}

//func summaryRanges(nums []int) []string {
//	if len(nums) == 0 {
//		return []string{}
//	}
//	if len(nums) == 1 {
//		return []string{fmt.Sprintf("%v", nums[0])}
//	}
//
//	res := make([]string, 0, len(nums))
//
//	j, i := 0, 0
//	for ; i < len(nums); i++ {
//		if nums[i]-nums[j] != i-j {
//			if i-j > 1 {
//				res = append(res, fmt.Sprintf("%v->%v", nums[j], nums[i-1]))
//			} else {
//				res = append(res, fmt.Sprintf("%v", nums[j]))
//			}
//			j = i
//		}
//	}
//	if i-j > 1 {
//		res = append(res, fmt.Sprintf("%v->%v", nums[j], nums[i-1]))
//	} else {
//		res = append(res, fmt.Sprintf("%v", nums[j]))
//	}
//	return res
//}
