package mid

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	notSatisfied := make([]int, n)

	sum := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 1 {
			notSatisfied[i] = customers[i]
		} else {
			sum += customers[i]
		}
	}

	reduce := 0
	for i := 0; i < minutes && i < n; i++ {
		reduce += notSatisfied[i]
	}
	maxReduce := reduce

	start := 0
	for end := minutes; end < n; end++ {
		reduce = reduce - notSatisfied[start] + notSatisfied[end]
		maxReduce = max(maxReduce, reduce)
		start++
	}

	return sum + maxReduce
}
