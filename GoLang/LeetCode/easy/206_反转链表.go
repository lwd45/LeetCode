package easy

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func reverseList(head *ListNode) *ListNode {
//	//if head == nil {
//	//	return nil
//	//}
//	//
//	//pre := ListNode{}
//	//for head != nil {
//	//	temp := head.Next
//	//	head.Next = pre.Next
//	//	pre.Next = head
//	//	head = temp
//	//}
//	//return pre.Next
//
//	var pre *ListNode
//	curr := head
//	for curr != nil {
//		next := curr.Next
//		curr.Next = pre
//		pre = curr
//		curr = next
//	}
//	return pre
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
