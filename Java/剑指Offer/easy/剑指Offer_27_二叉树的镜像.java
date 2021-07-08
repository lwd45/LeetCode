package easy;

import javax.swing.tree.TreeNode;

/**
 * Created by lwd at 2021/7/9
 *
 * @Description:
 */

public class 剑指Offer_27_二叉树的镜像 {
    public TreeNode mirrorTree(TreeNode root) {
        if (root == null) return null;

        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        mirrorTree(root.left);
        mirrorTree(root.right);
        return root;
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
