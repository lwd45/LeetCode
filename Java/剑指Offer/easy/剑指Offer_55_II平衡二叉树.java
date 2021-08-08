package easy;

/**
 * Created by lwd at 2021/8/8
 *
 * @Description:
 */
public class 剑指Offer_55_II平衡二叉树 {
    boolean balance = true;

    public boolean isBalanced(TreeNode root) {
        dfs(root);
        return balance;
    }

    private int dfs(TreeNode root) {
        if (root == null || !balance) return 0;
        int left = dfs(root.left) + 1;
        int right = dfs(root.right) + 1;
        balance = Math.abs(left - right) <= 1;
        return Math.max(left, right);
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
