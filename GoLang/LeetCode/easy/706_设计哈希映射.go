package easy

import "container/list"

const SIZE = 769

type entry struct {
	key, val int
}

type MyHashMap struct {
	elements []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{elements: make([]list.List, SIZE)}
}
func (this *MyHashMap) Hash(key int) int {
	return key % SIZE
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.Hash(key)

	for e := this.elements[idx].Front(); e != nil; e = e.Next() {
		if pair := e.Value.(entry); pair.key == key {
			e.Value = entry{key: key, val: value}
			return
		}
	}
	this.elements[idx].PushBack(entry{key: key, val: value})
}

func (this *MyHashMap) Get(key int) int {
	idx := this.Hash(key)

	for e := this.elements[idx].Front(); e != nil; e = e.Next() {
		if pair := e.Value.(entry); pair.key == key {
			return pair.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.Hash(key)

	for e := this.elements[idx].Front(); e != nil; e = e.Next() {
		if pair := e.Value.(entry); pair.key == key {
			this.elements[idx].Remove(e)
		}
	}
}
