package easy

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preNode := ListNode{
		Val:  0,
		Next: nil,
	}
	nowNode := &preNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			nowNode.Next = l1
			l1 = l1.Next
		} else {
			nowNode.Next = l2
			l2 = l2.Next
		}
		nowNode = nowNode.Next
	}
	if l1 != nil {
		nowNode.Next = l1
	} else {
		nowNode.Next = l2
	}
	return preNode.Next
}
