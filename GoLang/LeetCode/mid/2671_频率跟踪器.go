package mid

type FrequencyTracker struct {
	number map[int]int // key是number，val是该key的次数
	cnt    map[int]int // key是次数，val是还有多少不同的number是该次数
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		number: map[int]int{},
		cnt:    map[int]int{},
	}
}

func (this *FrequencyTracker) Add(number int) {
	oldCnt, _ := this.number[number]
	this.number[number]++

	if oldCnt > 0 {
		this.cnt[oldCnt]--
	}
	this.cnt[oldCnt+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	oldCnt, ok := this.number[number]

	if ok && oldCnt > 0 {
		this.number[number]--
		this.cnt[oldCnt]--
		if oldCnt-1> 0{
			this.cnt[oldCnt-1]++
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	frequency, ok := this.cnt[frequency]

	return ok && frequency > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
