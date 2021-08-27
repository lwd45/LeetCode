package mid;

import java.text.DateFormatSymbols;

/**
 * Created by lwd at 2021/8/27
 *
 * @Description:
 */
public class LeetCode_236_二叉树的最近公共祖先 {
    TreeNode ans = null;

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        dfs(root, p, q);
        return ans;
    }

    private boolean dfs(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || ans != null) return false;
        boolean left = dfs(root.left, p, q);
        boolean right = dfs(root.right, p, q);
        if ((left && right) || ((root.val == p.val || root.val == q.val) && (left || right)))
            ans = root;
        return left || right || root.val == p.val || root.val == q.val;
    }

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }
}
