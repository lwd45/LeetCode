package easy

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	ans := 0

	i, j := 0, len(nums)-1
	for ; i < j; {
		str := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		atoi, _ := strconv.Atoi(str)
		ans += atoi

		i++
		j--
	}

	if i == j {
		ans += nums[i]
	}

	return int64(ans)
}
