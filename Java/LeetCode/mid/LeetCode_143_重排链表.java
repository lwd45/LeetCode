package mid;

/**
 * Created by lwd at 2021/9/8
 *
 * @Description:
 */
public class LeetCode_143_重排链表 {
    public void reorderList(ListNode head) {
        if (head == null || head.next == null || head.next.next == null) return;
        ListNode slow = head, fast = head.next;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        ListNode l2 = slow.next;
        slow.next = null;
        l2 = reverse(l2);

        slow = head;
        while (l2 != null) {
            ListNode t = l2.next;
            l2.next = slow.next;
            slow.next = l2;
            slow = l2.next;
            l2 = t;
        }
    }

    public ListNode reverse(ListNode root) {
        if (root == null || root.next == null) return root;
        ListNode temp = reverse(root.next);
        root.next.next = root;
        root.next = null;
        return temp;
    }

    class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }
}
