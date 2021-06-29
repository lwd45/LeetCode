package mid

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	l1, l2 := head, slow.Next
	slow.Next = nil

	l2 = reverse(l2)
	for l2 != nil {
		temp := l2.Next
		l2.Next = l1.Next
		l1.Next = l2
		l1 = l1.Next.Next
		l2 = temp
	}
}

func reverse(l2 *ListNode) *ListNode {
	ans := &ListNode{}
	for l2 != nil {
		temp := l2.Next
		l2.Next = ans.Next
		ans.Next = l2
		l2 = temp
	}
	return ans.Next
}
