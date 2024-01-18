package mid

//func deleteDuplicates(head *ListNode) *ListNode {
//	prev := ListNode{}
//	prev.Next = head
//
//	now := &prev
//	for now.Next != nil && now.Next.Next != nil {
//		if now.Next.Val == now.Next.Next.Val {
//			val := now.Next.Val
//			next := now.Next.Next.Next
//			for next != nil && next.Val == val {
//				next = next.Next
//			}
//			now.Next = next
//		} else {
//			now = now.Next
//		}
//	}
//
//	return prev.Next
//}

func deleteDuplicates(head *ListNode) *ListNode {
	t := &ListNode{Next: head}

	fast, slow := t, head
	for slow != nil && slow.Next != nil {
		for slow.Next != nil && fast.Next.Val == slow.Next.Val {
			slow = slow.Next
		}

		if fast.Next != slow {
			fast.Next = slow.Next
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}

	return t.Next
}
