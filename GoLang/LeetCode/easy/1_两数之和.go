package easy

//func twoSum(nums []int, target int) []int {
//	mp := make(map[int][]int)
//	for i, num := range nums {
//		if _, ok := mp[num]; ok {
//			mp[num] = append(mp[num], i)
//		} else {
//			mp[num] = []int{i}
//		}
//	}
//
//	for i, num := range nums {
//		if value, ok := mp[target-num]; ok {
//			if target-num == num {
//				if len(value) > 1 {
//					return []int{value[0], value[1]}
//				}
//			} else {
//				return []int{i, value[0]}
//			}
//		}
//	}
//	return []int{-1, -1}
//}

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	for idx, num := range nums {
		if _, exist := idxMap[target-num]; exist {
			return []int{idxMap[target-num], idx}
		}
		idxMap[num] = idx
	}
	return nil
}
