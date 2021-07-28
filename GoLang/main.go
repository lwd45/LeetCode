package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	digit := 1 //位数
	count := 9 //位数
	start := 1 //起始

	for n > count {
		n -= count
		digit++
		start *= 10
		count = digit * 9 * start
	}

	num := start + (n-1)/digit
	str := strconv.Itoa(num)
	return int(str[(n-1)%digit] - '0')
}
func main() {
	fmt.Println(findNthDigit(100))
}
