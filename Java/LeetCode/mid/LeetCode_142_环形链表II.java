package mid;

/**
 * Created by lwd at 2021/9/2
 *
 * @Description:
 */
public class LeetCode_142_环形链表II {
    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null || head.next.next == null) return null;

        ListNode fast = head.next.next, slow = head.next;
        while (fast != null && fast.next != null && fast != slow) {
            fast = fast.next.next;
            slow = slow.next;
        }
        if (fast == null || fast != slow) {
            return null;
        }
        slow = head;
        while (slow != fast) {
            slow = slow.next;
            fast = fast.next;
        }
        return slow;
    }

    class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
            next = null;
        }
    }
}
