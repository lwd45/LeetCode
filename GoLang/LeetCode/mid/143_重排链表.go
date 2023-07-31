package mid

//func reorderList(head *ListNode) {
//	if head == nil || head.Next == nil || head.Next.Next == nil {
//		return
//	}
//	fast, slow := head.Next, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//
//	l1, l2 := head, slow.Next
//	slow.Next = nil
//
//	l2 = reverse(l2)
//	for l2 != nil {
//		temp := l2.Next
//		l2.Next = l1.Next
//		l1.Next = l2
//		l1 = l1.Next.Next
//		l2 = temp
//	}
//}
//
//func reverse(l2 *ListNode) *ListNode {
//	ans := &ListNode{}
//	for l2 != nil {
//		temp := l2.Next
//		l2.Next = ans.Next
//		ans.Next = l2
//		l2 = temp
//	}
//	return ans.Next
//}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil

	l1, l2 := head, reverse(fast)

	for l2 != nil {
		t := l1.Next
		l1.Next = l2
		t2 := l2.Next
		l2.Next = t
		l2 = t2
		l1 = l1.Next.Next
	}
}

func reverse(fast *ListNode) *ListNode {
	temp := &ListNode{}
	for fast != nil {
		t := fast.Next
		fast.Next = temp.Next
		temp.Next = fast
		fast = t
	}
	return temp.Next
}
