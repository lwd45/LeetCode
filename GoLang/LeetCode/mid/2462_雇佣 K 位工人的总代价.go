package mid

import (
	"container/heap"
)

type intHeap [][]int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	if (*h)[i][0] == (*h)[j][0] {
		return (*h)[i][1] <= (*h)[j][1]
	}
	return (*h)[i][0] < (*h)[j][0]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *intHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	hp := &intHeap{}

	left, right := candidates-1, n-candidates
	if left+1 < right {
		for i := 0; i <= left; i++ {
			heap.Push(hp, []int{costs[i], i})
		}
		for i := right; i < n; i++ {
			heap.Push(hp, []int{costs[i], i})
		}
	} else {
		for i := 0; i < n; i++ {
			heap.Push(hp, []int{costs[i], i})
		}
	}

	cost := 0
	for i := 0; i < k; i++ {
		pop := heap.Pop(hp).([]int)
		cost += pop[0]

		if left+1 < right {
			if pop[1] > left {
				right--
				heap.Push(hp, []int{costs[right], right})
			} else {
				left++
				heap.Push(hp, []int{costs[left], left})
			}
		}
	}
	return int64(cost)
}
