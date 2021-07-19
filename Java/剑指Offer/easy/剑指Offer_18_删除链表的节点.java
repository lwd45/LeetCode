package easy;

import java.util.List;

/**
 * Created by lwd at 2021/7/19
 *
 * @Description:
 */
public class 剑指Offer_18_删除链表的节点 {
    public ListNode deleteNode(ListNode head, int val) {
        ListNode pre = new ListNode(-1);
        pre.next = head;
        ListNode now = pre;

        while (now.next != null) {
            if (now.next.val == val) {
                now.next = now.next.next;
                return pre.next;
            }
            now = now.next;
        }
        return pre.next;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}