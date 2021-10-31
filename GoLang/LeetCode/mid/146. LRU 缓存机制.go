package mid

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	length   int
	first    *Node
	last     *Node
	maps     *map[int]*Node
}

func Constructors(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		length:   0,
		first:    &Node{},
		last:     &Node{},
		maps:     &map[int]*Node{},
	}
	cache.first.next = cache.last
	cache.last.pre = cache.first
	return cache
}

func (this *LRUCache) Get(key int) int {
	if value, ok := (*this.maps)[key]; ok {
		value.next.pre = value.pre
		value.pre.next = value.next
		this.moveToHead(value)
		return value.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := (*this.maps)[key]; ok {
		node.next.pre = node.pre
		node.pre.next = node.next
		node.value = value
		this.moveToHead(node)
	} else {
		if this.length < this.capacity {
			this.length++
		} else {
			last := this.removeLast()
			delete(*this.maps, last.key)
		}
		n := &Node{
			key:   key,
			value: value,
			pre:   this.first,
			next:  this.first.next,
		}
		this.moveToHead(n)
		(*this.maps)[key] = n
	}
}

func (this *LRUCache) removeLast() *Node {
	temp := this.last.pre
	temp.pre.next = temp.next
	temp.next.pre = temp.pre
	return temp
}

func (this *LRUCache) moveToHead(node *Node) {
	this.first.next.pre = node
	node.next = this.first.next
	node.pre = this.first
	this.first.next = node
}
