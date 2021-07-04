package easy

//func findErrorNums(nums []int) (ans []int) {
//	visited := make([]bool, len(nums)+1)
//	for i := range nums {
//		if visited[nums[i]] {
//			ans = append(ans, nums[i])
//		} else {
//			visited[nums[i]] = true
//		}
//	}
//	for i := 1; i <= len(visited); i++ {
//		if !visited[i] {
//			ans = append(ans, i)
//		}
//	}
//	return ans
//}

func findErrorNums(nums []int) (ans []int) {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			swap(&nums, i, nums[i]-1)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, nums[i], i+1)
			return
		}
	}
	return
}

func swap(nums *[]int, idx1, idx2 int) {
	temp := (*nums)[idx1]
	(*nums)[idx1] = (*nums)[idx2]
	(*nums)[idx2] = temp
}
