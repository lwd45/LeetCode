package easy;

/**
 * Created by lwd at 2021/7/16
 *
 * @Description:
 */
public class 剑指Offer_24_反转链表 {
    // public ListNode reverseList(ListNode head) {
    //     ListNode preHead = new ListNode(-1);
    //
    //     while (head != null) {
    //         ListNode temp = head.next;
    //         head.next = preHead.next;
    //         preHead.next = head;
    //         head = temp;
    //     }
    //     return preHead.next;
    // }

    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode node = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return node;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}