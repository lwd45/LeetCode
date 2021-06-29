package easy

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	prevA, prevB := &ListNode{}, &ListNode{}
	prevA.Next, prevB.Next = headA, headB

	nowA, nowB := prevA, prevB
	lenA, lenB := 0, 0
	for nowA.Next != nil {
		lenA++
		nowA = nowA.Next
	}
	for nowB.Next != nil {
		lenB++
		nowB = nowB.Next
	}

	for lenA > lenB {
		lenA--
		prevA = prevA.Next
	}
	for lenB > lenA {
		lenB--
		prevB = prevB.Next
	}

	for prevA != nil && prevB != nil {
		if prevA == prevB {
			return prevA
		}
		prevA = prevA.Next
		prevB = prevB.Next
	}
	return nil
}
