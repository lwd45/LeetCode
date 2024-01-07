package mid

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	temp := head
	for temp.Next != nil {
		big, small := max(temp.Val, temp.Next.Val), min(temp.Val, temp.Next.Val)
		num := gcd(big, small)

		insert := &ListNode{Val: num, Next: temp.Next}
		temp.Next = insert
		temp = temp.Next.Next
	}
	return head
}

func gcd(v1, v2 int) int {
	if v1%v2 == 0 {
		return v2
	}
	return gcd(v2, v1%v2)
}
