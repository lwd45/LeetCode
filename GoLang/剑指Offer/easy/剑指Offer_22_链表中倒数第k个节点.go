package easy

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
