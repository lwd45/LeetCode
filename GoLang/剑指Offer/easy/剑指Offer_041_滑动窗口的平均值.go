package easy

type MovingAverage struct {
	slice []int
	size int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		slice: make([]int, 0, size),
		size:   size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.slice) >= this.size {
		this.sum -= this.slice[0]
		this.slice = this.slice[1:]
	}
	this.slice = append(this.slice, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.slice))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
