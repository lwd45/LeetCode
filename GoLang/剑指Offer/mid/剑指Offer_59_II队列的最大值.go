package mid

type MaxQueue struct {
	queueSlice []int
	maxSlice   []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queueSlice: make([]int, 0),
		maxSlice:   make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxSlice) == 0 {
		return -1
	}
	return this.maxSlice[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queueSlice = append(this.queueSlice, value)
	for len(this.maxSlice) > 0 && this.maxSlice[len(this.maxSlice)-1] < value {
		this.maxSlice = this.maxSlice[:len(this.maxSlice)-1]
	}
	this.maxSlice = append(this.maxSlice, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queueSlice) == 0 {
		return -1
	}
	pop := this.queueSlice[0]
	this.queueSlice = this.queueSlice[1:]
	if pop == this.maxSlice[0] {
		this.maxSlice = this.maxSlice[1:]
	}
	return pop
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
