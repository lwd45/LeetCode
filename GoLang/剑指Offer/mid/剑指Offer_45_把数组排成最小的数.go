package mid

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	slic := make([]string, len(nums))
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
