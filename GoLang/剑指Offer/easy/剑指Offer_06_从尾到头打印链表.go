package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	reverseHelper(head, &ans)
	return ans
}

func reverseHelper(head *ListNode, ans *[]int) {
	if head.Next != nil {
		reverseHelper(head.Next, ans)
	}
	*ans = append(*ans, head.Val)
}
