package easy;

/**
 * Created by lwd at 2021/8/6
 *
 * @Description:
 */
public class 剑指Offer_54_二叉搜索树的第k大节点 {
    private int index;
    private int ans;

    public int kthLargest(TreeNode root, int k) {
        index = k;
        preOrder(root);
        return ans;
    }

    private void preOrder(TreeNode root) {
        if (root == null || index < 0) {
            return;
        }

        preOrder(root.right);
        index--;
        if (index == 0) {
            ans = root.val;
        }
        preOrder(root.left);
    }

    class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

}
