package mid

import "strconv"

func findNthDigit(n int) int {
	digit := 1 //位数
	count := 9 //位数
	base := 9
	start := 1 //起始

	for n > count {
		digit++
		base *= 10
		count += base
		start *= 10
	}

	num := start + (n-1)/digit
	str := strconv.Itoa(num)
	return int(str[(n-1)%digit])
}