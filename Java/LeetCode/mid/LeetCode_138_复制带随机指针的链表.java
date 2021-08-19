package mid;

/**
 * Created by lwd at 2021/8/19
 *
 * @Description:
 */
public class LeetCode_138_复制带随机指针的链表 {
    public Node copyRandomList(Node head) {
        if (head == null) return null;

        Node temp = head;
        while (temp != null) {
            Node node = new Node(temp.val);
            node.next = temp.next;
            temp.next = node;
            temp = temp.next.next;
        }

        temp = head;
        while (temp != null) {
            if (temp.random != null)
                temp.next.random = temp.random.next;
            temp = temp.next.next;
        }

        Node origin = new Node(-1);
        Node ans = new Node(-1);
        origin.next = head;
        ans.next = head.next;

        Node temp1 = head;
        temp = head.next;
        while (temp1 != null) {
            temp1.next = temp1.next.next;
            temp.next = temp.next == null ? null : temp.next.next;

            temp1 = temp1.next;
            temp = temp.next;
        }
        return ans.next;
    }

    class Node {
        int val;
        Node next;
        Node random;

        public Node(int val) {
            this.val = val;
            this.next = null;
            this.random = null;
        }
    }
}
