package easy;

/**
 * Created by lwd at 2021/8/10
 *
 * @Description:
 */
public class 剑指Offer_68_II二叉树的最近公共祖先 {
    private TreeNode ans;

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        ans = null;
        dfs(root, p, q);
        return ans;
    }

    private boolean dfs(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || ans != null) return false;
        boolean lSon = dfs(root.left, p, q);
        boolean rSon = dfs(root.right, p, q);
        if ((lSon && rSon) || ((lSon || rSon) && (root.val == q.val || root.val == p.val))) {
            ans = root;
        }
        return (lSon || rSon || (root.val == p.val) || (root.val == q.val));
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
