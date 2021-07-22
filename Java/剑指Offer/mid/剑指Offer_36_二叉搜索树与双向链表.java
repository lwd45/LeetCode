package mid;

/**
 * Created by lwd at 2021/7/23
 *
 * @Description:
 */
public class 剑指Offer_36_二叉搜索树与双向链表 {
    Node pre, head;

    public Node treeToDoublyList(Node root) {
        if (root == null) return root;
        dfs(root);

        head.left = pre;
        pre.right = head;
        return head;
    }

    private void dfs(Node root) {
        if (root == null) return;
        dfs(root.left);
        if (pre != null) {
            pre.right = root;
            root.left = pre;
        } else {
            head = root;
        }
        pre = root;
        dfs(root.right);
    }

    class Node {
        public int val;
        public Node left;
        public Node right;

        public Node() {
        }

        public Node(int _val) {
            val = _val;
        }

        public Node(int _val, Node _left, Node _right) {
            val = _val;
            left = _left;
            right = _right;
        }
    }
}
