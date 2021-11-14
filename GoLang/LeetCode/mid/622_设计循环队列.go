package mid
//
//type MyCircularQueue struct {
//	capacity    int
//	length      int
//	insertIndex int
//	outputIndex int
//	array       []int
//}
//
//func Constructor(k int) MyCircularQueue {
//	return MyCircularQueue{
//		capacity:    k,
//		length:      0,
//		insertIndex: -1,
//		outputIndex: 0,
//		array:       make([]int, k),
//	}
//}
//
//func (this *MyCircularQueue) EnQueue(value int) bool {
//	if this.length < this.capacity {
//		this.length++
//		this.insertIndex = (this.insertIndex + 1) % this.capacity
//		this.array[this.insertIndex] = value
//		return true
//	}
//	return false
//}
//
//func (this *MyCircularQueue) DeQueue() bool {
//	if this.length > 0 {
//		this.length--
//		this.outputIndex = (this.outputIndex + 1) % this.capacity
//		return true
//	}
//	return false
//}
//
//func (this *MyCircularQueue) Front() int {
//	if this.length > 0 {
//		return this.array[this.outputIndex]
//	}
//	return -1
//}
//
//func (this *MyCircularQueue) Rear() int {
//	if this.length > 0 {
//		return this.array[this.insertIndex]
//	}
//	return -1
//}
//
//func (this *MyCircularQueue) IsEmpty() bool {
//	return this.length == 0
//}
//
//func (this *MyCircularQueue) IsFull() bool {
//	return this.length == this.capacity
//}
