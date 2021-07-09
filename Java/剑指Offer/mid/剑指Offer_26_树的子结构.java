package mid;

import javax.swing.tree.TreeNode;

/**
 * Created by lwd at 2021/7/9
 *
 * @Description:
 */
public class 剑指Offer_26_树的子结构 {

    public boolean isSubStructure(TreeNode A, TreeNode B) {
        if (A == null || B == null) return false;
        return isSubStructure(A.left, B) || isSubStructure(A.right, B) || curSub(A, B);
    }

    private boolean curSub(TreeNode A, TreeNode B) {
        if (B == null) return true;
        if (A == null || A.val != B.val) return false;
        else return curSub(A.left, B.left) && curSub(A.right, B.right);
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
