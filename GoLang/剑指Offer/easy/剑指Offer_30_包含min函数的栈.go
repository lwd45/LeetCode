package easy

type MinStack struct {
	stackNormal []int
	stackMin    []int
}

func Constructor() MinStack {
	return MinStack{
		stackNormal: make([]int, 0),
		stackMin:    make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stackNormal = append(this.stackNormal, x)
	if len(this.stackMin) == 0 || this.stackMin[len(this.stackMin)-1] >= x {
		this.stackMin = append(this.stackMin, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stackNormal) == 0 {
		return
	}
	pop := this.stackNormal[len(this.stackNormal)-1]
	this.stackNormal = this.stackNormal[:len(this.stackNormal)-1]
	if this.stackMin[len(this.stackMin)-1] == pop {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stackNormal[len(this.stackNormal)-1]
}

func (this *MinStack) Min() int {
	return this.stackMin[len(this.stackMin)-1]
}
