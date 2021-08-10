package easy;

/**
 * Created by lwd at 2021/8/10
 *
 * @Description:
 */
public class 剑指Offer_68_I二叉搜索树的最近公共祖先 {

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        while (true) {
            if (root.val > p.val && root.val > q.val) return lowestCommonAncestor(root.left, p, q);
            else if (root.val < p.val && root.val < q.val) return lowestCommonAncestor(root.right, p, q);
            else return root;
        }
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
