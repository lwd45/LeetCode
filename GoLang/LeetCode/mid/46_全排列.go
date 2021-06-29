package mid

func permute(nums []int) [][]int {
	oneList := make([]int, 0, len(nums))
	for _, v := range nums {
		oneList = append(oneList, v)
	}

	var ans [][]int
	dfs46(&ans, &oneList, 0, len(nums))
	return ans
}

func dfs46(ans *[][]int, oneList *[]int, index int, length int) {
	if index == length {
		var copyList = make([]int, len(*oneList))
		copy(copyList, *oneList)
		*ans = append(*ans, copyList)
	}

	for i := index; i < length; i++ {
		swap46(oneList, i, index)
		dfs46(ans, oneList, index+1, length)
		swap46(oneList, i, index)
	}
}

func swap46(nums *[]int, i, j int) {
	temp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}
