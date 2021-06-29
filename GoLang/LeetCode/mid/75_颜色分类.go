package mid

import "fmt"

func sortColors(nums []int) {
	L, R := 0, len(nums)-1
	for index := 0; index <= R; {
		if nums[index] == 0 {
			swap75(&nums, L, index)
			L++
			index++
		} else if nums[index] == 2 {
			swap75(&nums, R, index)
			R--
		} else {
			index++
		}
	}
	fmt.Println(nums)
}

func swap75(nums *[]int, i, j int) {
	temp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}
