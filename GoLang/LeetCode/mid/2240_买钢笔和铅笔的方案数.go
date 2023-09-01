package mid

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if total < cost1 && total < cost2 {
		return 1
	}

	var ans int64
	max1 := total / cost1
	for i := 0; i <= max1; i++ {
		ans += int64((total-i*cost1)/cost2) + 1
	}
	return ans
}
