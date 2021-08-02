package easy;

import java.util.ArrayList;
import java.util.LinkedList;

/**
 * Created by lwd at 2021/8/2
 *
 * @Description:
 */
public class 剑指Offer_55_I二叉树的深度 {
    // public int maxDepth(TreeNode root) {
    //     if (root == null) return 0;
    //     int deep = 0;
    //     LinkedList<TreeNode> list = new LinkedList<>();
    //     list.addLast(root);
    //     while (list.size() > 0) {
    //         deep++;
    //         int size = list.size();
    //         for (int i = size; i > 0; i--) {
    //             TreeNode pop = list.pop();
    //             if (pop.left != null) {
    //                 list.addLast(pop.left);
    //             }
    //             if (pop.left != null) {
    //                 list.addLast(pop.right);
    //             }
    //         }
    //     }
    //     return deep;
    // }

    public int maxDepth(TreeNode root) {
        return help(root, 0);
    }

    private int help(TreeNode root, int deep) {
        if (root == null) return deep;

        int left = help(root.left, deep + 1);
        int right = help(root.right, deep + 1);
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
