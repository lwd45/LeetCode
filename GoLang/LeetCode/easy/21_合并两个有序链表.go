package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev = ListNode{}
	var now = &prev

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			now.Next = l2
			l2 = l2.Next
		} else {
			now.Next = l1
			l1 = l1.Next
		}
		now = now.Next
	}

	if l1 != nil {
		now.Next = l1
	}

	if l2 != nil {
		now.Next = l2
	}

	return prev.Next
}
