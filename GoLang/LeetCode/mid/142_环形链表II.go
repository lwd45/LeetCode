package mid

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil || fast.Next != slow {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
