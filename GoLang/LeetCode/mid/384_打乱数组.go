package mid

import (
	"math/rand"
)

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();

给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
实现 Solution class:
Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果

输入
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
输出
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

提示：
1 <= nums.length <= 200
-106 <= nums[i] <= 106
nums 中的所有元素都是唯一的
最多可以调用 5 * 104 次 reset 和 shuffle
*/

type Solution struct {
	origin  []int
	shuffle []int
}

//func Constructor(nums []int) Solution {
//	return Solution{origin: append([]int{}, nums...), shuffle: nums}
//}

func (this *Solution) Reset() []int {
	return this.origin
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.shuffle); i++ {
		idx := rand.Intn(len(this.shuffle)-i) + i
		t := this.shuffle[i]
		this.shuffle[i] = this.shuffle[idx]
		this.shuffle[idx] = t
	}
	return this.shuffle
}
