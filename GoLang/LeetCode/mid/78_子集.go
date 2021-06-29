package mid

func subsets(nums []int) [][]int {
	var ans [][]int
	var one []int
	dfs78(nums, &one, &ans, 0)
	return ans
}

func dfs78(nums []int, one *[]int, ans *[][]int, i int) {
	if i > len(nums) {
		return
	}

	//这里一定要拷贝
	oneAns := make([]int, len(*one))
	copy(oneAns, *one)
	*ans = append(*ans, oneAns)

	for index := i; index < len(nums); index++ {
		*one = append(*one, nums[index])
		dfs78(nums, one, ans, index+1)
		*one = (*one)[:len(*one)-1]
	}
}
