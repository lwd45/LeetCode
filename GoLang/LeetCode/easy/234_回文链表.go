package easy

//func IsPalindrome(head *ListNode) bool {
//	prev := ListNode{}
//	prev.Next = nil
//
//	now := head
//	for now != nil {
//		node := &ListNode{
//			Val:  now.Val,
//			Next: prev.Next,
//		}
//		prev.Next = node
//		now = now.Next
//	}
//
//	second := prev.Next
//	for head != nil {
//		if head.Val != second.Val {
//			return false
//		}
//		head = head.Next
//		second = second.Next
//	}
//	return true
//}

//func isPalindrome(head *ListNode) bool {
//	var vals []int
//	for ; head != nil; head = head.Next {
//		vals = append(vals, head.Val)
//	}
//
//	i := 0
//	j := len(vals) - 1
//	for i < j {
//		if vals[i] != vals[j] {
//			return false
//		} else {
//			i++
//			j--
//		}
//	}
//	return true
//}
