package hard

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{}
	pre.Next = nil

	faster, slow, now := head, head, pre
	for faster != nil {
		i := k - 1
		for faster != nil && i > 0 {
			faster = faster.Next
			i--
		}
		if faster != nil {
			nowStart, nextStart := slow, faster.Next
			for slow != faster {
				t := slow.Next
				slow.Next = now.Next
				now.Next = slow
				slow = t
			}
			slow.Next = now.Next
			now.Next = slow
			now = nowStart
			slow, faster = nextStart, nextStart
		} else {
			now.Next = slow
		}
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
