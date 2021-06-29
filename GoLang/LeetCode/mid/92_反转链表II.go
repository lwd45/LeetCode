package mid

/*
给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]


链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	prev := &ListNode{}
	prev.Next = head

	now := prev
	for i := 1; i < left; i++ {
		now = now.Next
	}

	now.Next = reverseNode(now.Next, right-left+1)
	return prev.Next
}

func reverseNode(next *ListNode, k int) *ListNode {
	prev := &ListNode{}
	start := next
	for ; k > 0; k-- {
		temp := next.Next
		next.Next = prev.Next
		prev.Next = next
		next = temp
	}
	start.Next = next
	return prev.Next
}
