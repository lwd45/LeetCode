package mid

func deleteDuplicates(head *ListNode) *ListNode {
	prev := ListNode{}
	prev.Next = head

	now := &prev
	for now.Next != nil && now.Next.Next != nil {
		if now.Next.Val == now.Next.Next.Val {
			val := now.Next.Val
			next := now.Next.Next.Next
			for next != nil && next.Val == val {
				next = next.Next
			}
			now.Next = next
		} else {
			now = now.Next
		}
	}

	return prev.Next
}
