package mid

import (
	"math/rand"
)

//type PriorityQueue struct {
//	sort.IntSlice
//}
//
//func (p *PriorityQueue) Less(i, j int) bool {
//	return p.IntSlice[i] > p.IntSlice[j]
//}
//
//func (p *PriorityQueue) Push(x interface{}) {
//	p.IntSlice = append(p.IntSlice, x.(int))
//}
//
//func (p *PriorityQueue) Pop() interface{} {
//	m := len(p.IntSlice)
//	v := p.IntSlice[m-1]
//	p.IntSlice = p.IntSlice[:m-1]
//	return v
//}
//
//func findKthLargest(nums []int, k int) int {
//	q := &PriorityQueue{nums}
//	heap.Init(q)
//
//	v := 0
//	for i := 0; i < k; i++ {
//		v = heap.Pop(q).(int)
//	}
//	return v
//}

func findKthLargest(nums []int, k int) int {
	return quickSort_215(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort_215(nums []int, l, r, k int) int {
	idx := sort_215(nums, l, r)
	if idx == k {
		return nums[idx]
	} else if idx > k {
		return quickSort_215(nums, l, idx-1, k)
	} else {
		return quickSort_215(nums, idx+1, r, k)
	}
}

func sort_215(nums []int, l, r int) int {
	if l == r {
		return l
	}

	idx := rand.Intn(r-l+1) + l

	compare := nums[idx]
	nums[idx] = nums[r]
	left, right := l, r-1
	for left <= right {
		for left <= right && nums[left] < compare {
			left++
		}
		for left <= right && nums[right] >= compare {
			right--
		}
		if left < right {
			t := nums[left]
			nums[left] = nums[right]
			nums[right] = t
		}
	}
	nums[r] = nums[left]
	nums[left] = compare
	return left
}
