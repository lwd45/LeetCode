package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//fmt.Println(strconv.ParseInt("09", 10, 64))
	fmt.Println(minNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func minNumber(nums []int) string {
	slic := make([]string, 0)
	for _, v := range nums {
		slic = append(slic, strconv.Itoa(v))
	}
	sort.Slice(slic, func(i, j int) bool {
		num1, _ := strconv.ParseInt(slic[i]+slic[j], 10, 64)
		num2, _ := strconv.ParseInt(slic[j]+slic[i], 10, 64)
		return num1 < num2
	})
	ans := ""
	for _, v := range slic {
		ans += v
	}
	return ans
}
