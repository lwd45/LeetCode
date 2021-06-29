package mid

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(&nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums *[]int, start, end int) {
	if start < end {
		pos := quickSortOnce(nums, start, end)
		quickSort(nums, start, pos-1)
		quickSort(nums, pos+1, end)
	}
}

func quickSortOnce(nums *[]int, start, end int) int {
	pos := start + rand.Intn(end-start+1)
	swap912(nums, end, pos)
	comp := (*nums)[end]
	index := start
	for i := start; i < end; i++ {
		if (*nums)[i] < comp {
			swap912(nums, i, index)
			index++
		}
	}
	swap912(nums, end, index)
	return index
}

func swap912(nums *[]int, i, j int) {
	temp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}
