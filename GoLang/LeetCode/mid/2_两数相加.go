package mid

//func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
//	prev := &ListNode{}
//	now := prev
//
//	carry := 0
//	for l1 != nil || l2 != nil {
//		num1, num2, sum := 0, 0, 0
//		if l1 != nil {
//			num1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			num2 = l2.Val
//			l2 = l2.Next
//		}
//		sum = num1 + num2 + carry
//		carry = sum / 10
//		sum = sum % 10
//
//		now.Next = &ListNode{Val: sum}
//		now = now.Next
//	}
//
//	if carry == 1 {
//		now.Next = &ListNode{Val: 1}
//	}
//	return prev.Next
//}

//func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
//	prev := &ListNode{}
//	now := prev
//
//	carry := 0
//	for l1 != nil && l2 != nil {
//		sum := l1.Val + l2.Val + carry
//		carry = sum / 10
//		sum = sum % 10
//
//		now.Next = &ListNode{Val: sum}
//		now = now.Next
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//
//	if l1 != nil {
//		if carry == 0 {
//			now.Next = l1
//		} else {
//			for l1 != nil && l1.Val == 9 {
//				now.Next = &ListNode{Val: 0}
//				now = now.Next
//				l1 = l1.Next
//			}
//			if l1 != nil {
//				now.Next = l1
//				l1.Val++
//				carry = 0
//			}
//		}
//	}
//
//	if l2 != nil {
//		if carry == 0 {
//			now.Next = l2
//		} else {
//			for l2 != nil && l2.Val == 9 {
//				now.Next = &ListNode{Val: 0}
//				now = now.Next
//				l2 = l2.Next
//			}
//			if l2 != nil {
//				now.Next = l2
//				l2.Val++
//				carry = 0
//			}
//		}
//	}
//
//	if carry == 1 {
//		now.Next = &ListNode{Val: 1}
//	}
//
//	return prev.Next
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	now := temp

	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		val := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10

		node := &ListNode{
			Val: val,
		}
		now.Next = node
		now = now.Next
	}

	for ; l1 != nil; l1 = l1.Next {
		val := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10

		node := &ListNode{
			Val: val,
		}
		now.Next = node
		now = now.Next
	}

	for ; l2 != nil; l2 = l2.Next {
		val := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10

		node := &ListNode{
			Val: val,
		}
		now.Next = node
		now = now.Next
	}

	if carry > 0 {
		node := &ListNode{
			Val: carry,
		}
		now.Next = node
		now = now.Next
	}

	return temp.Next
}
