package mid

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	temp, length := head, 0
	for temp != nil {
		length++
		temp = temp.Next
	}

	k = k % length
	if k == 0 {
		return head
	}
	prev := &ListNode{}
	prev.Next = head
	slow, fast := prev, prev
	for k > 0 {
		k--
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	ans := slow.Next
	slow.Next, fast.Next = nil, head
	return ans
}
