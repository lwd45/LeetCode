package mid

import (
	"container/heap"
	"sort"
)

type PriorityQueue struct {
	sort.IntSlice
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p.IntSlice[i] > p.IntSlice[j]
}

func (p *PriorityQueue) Push(v interface{}) {
	p.IntSlice = append(p.IntSlice, v.(int))
}

func (p *PriorityQueue) Pop() interface{} {
	old := p.IntSlice
	l := len(old)
	v := old[l-1]
	p.IntSlice = p.IntSlice[:l-1]
	return v
}

func minStoneSum(piles []int, k int) int {
	pq := &PriorityQueue{piles}
	heap.Init(pq)
	for i := 0; i < k; i++ {
		pile := heap.Pop(pq).(int)
		pile -= pile / 2
		heap.Push(pq, pile)
	}

	sum := 0
	for len(pq.IntSlice) > 0 {
		sum += heap.Pop(pq).(int)
	}
	return sum
}
