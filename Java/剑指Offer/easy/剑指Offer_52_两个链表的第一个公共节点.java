package easy;

import java.util.List;

/**
 * Created by lwd at 2021/7/21
 *
 * @Description:
 */
public class 剑指Offer_52_两个链表的第一个公共节点 {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) return null;

        int len1 = getLen(headA), len2 = getLen(headB);
        if (len2 > len1) {
            ListNode temp = headA;
            headA = headB;
            headB = temp;

            int tempN = len1;
            len1 = len2;
            len2 = tempN;
        }
        while (len1 > len2) {
            headA = headA.next;
            len1--;
        }

        while (headA != null && headA != headB) {
            headA = headA.next;
            headB = headB.next;
        }
        return headA;
    }

    private int getLen(ListNode node) {
        ListNode temp = node;
        int ans = 0;
        while (temp != null) {
            ans++;
            temp = temp.next;
        }
        return ans;
    }

}


