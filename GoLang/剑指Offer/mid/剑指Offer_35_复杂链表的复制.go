package mid

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	oldHead := head
	for oldHead != nil {
		newNode := &Node{Val: oldHead.Val, Next: oldHead.Next}
		oldHead.Next = newNode
		oldHead = oldHead.Next.Next
	}

	oldHead = head
	for oldHead != nil {
		if oldHead.Random != nil {
			oldHead.Next.Random = oldHead.Random.Next
		}
		oldHead = oldHead.Next.Next
	}

	first := head
	second := head.Next
	ans := second
	for first != nil {
		first.Next = first.Next.Next
		first = first.Next

		if first != nil {
			second.Next = second.Next.Next
			second = second.Next
		}
	}
	return ans
}
