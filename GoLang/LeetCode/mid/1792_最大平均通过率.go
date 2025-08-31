package mid

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := hp{}
	for _, class := range classes {
		a, b := class[0], class[1]
		x := float64(a+1)/float64(b+1) - float64(a)/float64(b)
		heap.Push(&h, tuple{x, b, a})
	}
	for i := 0; i < extraStudents; i++ {
		e := heap.Pop(&h).(tuple)
		a, b := e.a+1, e.b+1
		x := float64(a+1)/float64(b+1) - float64(a)/float64(b)
		heap.Push(&h, tuple{x, b, a})
	}

	var ans float64
	for len(h) > 0 {
		e := heap.Pop(&h).(tuple)
		ans += float64(e.a) / float64(e.b)
	}
	return ans / float64(len(classes))
}

type tuple struct {
	x    float64
	a, b int
}

type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.x > b.x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { v := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return v }
