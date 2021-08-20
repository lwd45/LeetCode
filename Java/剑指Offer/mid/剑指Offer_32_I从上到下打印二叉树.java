package mid;

import java.util.ArrayList;
import java.util.LinkedList;

/**
 * Created by lwd at 2021/7/26
 *
 * @Description:
 */
public class 剑指Offer_32_I从上到下打印二叉树 {
    public int[] levelOrder(TreeNode root) {
        if (root == null) return new int[0];

        ArrayList<TreeNode> ans = new ArrayList<>();
        LinkedList<TreeNode> list = new LinkedList<>();
        list.addLast(root);
        ans.add(root);
        while (!list.isEmpty()) {
            TreeNode first = list.pollFirst();
            if (first.left != null) {
                list.addLast(first.left);
                ans.add(first.left);
            }
            if (first.right != null) {
                list.addLast(first.right);
                ans.add(first.right);
            }
        }
        int[] ret = new int[ans.size()];
        for (int i = 0; i < ret.length; ++i) {
            ret[i] = ans.get(i).val;
        }
        return ret;
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
