package mid
//
////import "strings"
////
////type MapSum struct {
////	cnt map[string]int
////}
////
////func Constructor() MapSum {
////	return MapSum{cnt: map[string]int{}}
////}
////
////func (this *MapSum) Insert(key string, val int) {
////	this.cnt[key] = val
////}
////
////func (this *MapSum) Sum(prefix string) int {
////	sum := 0
////	for key, val := range this.cnt {
////		if strings.HasPrefix(key, prefix) {
////			sum += val
////		}
////	}
////	return sum
////}
//
////type MapSum struct {
////	cnt map[string]int
////	pre map[string]int
////}
////
////func Constructor() MapSum {
////	return MapSum{
////		cnt: map[string]int{},
////		pre: map[string]int{},
////	}
////}
////
////func (this *MapSum) Insert(key string, val int) {
////	delta := val
////	if old, ok := this.cnt[key]; ok {
////		delta = val - old
////	}
////	this.cnt[key] = val
////	for i := range key {
////		this.pre[key[:i+1]] = this.pre[key[:i+1]] + delta
////	}
////}
////
////func (this *MapSum) Sum(prefix string) int {
////	return this.pre[prefix]
////}
//
//type TrieNode struct {
//	children [26]*TrieNode
//	val      int
//}
//
//type MapSum struct {
//	root *TrieNode
//	cnt  map[string]int
//}
//
//func Constructor() MapSum {
//	return MapSum{
//		root: &TrieNode{},
//		cnt:  map[string]int{},
//	}
//}
//
//func (this *MapSum) Insert(key string, val int) {
//	delta := val
//	if old, ok := this.cnt[key]; ok {
//		delta -= old
//	}
//	this.cnt[key] = val
//	root := this.root
//	for _, c := range key {
//		c -= 'a'
//		if root.children[c] == nil {
//			root.children[c] = &TrieNode{}
//		}
//		root = root.children[c]
//		root.val += delta
//	}
//}
//
//func (this *MapSum) Sum(prefix string) int {
//	root := this.root
//	for _, c := range prefix {
//		c = c - 'a'
//		if root.children[c] == nil {
//			return 0
//		}
//		root = root.children[c]
//	}
//	return root.val
//}
