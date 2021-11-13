package mid;

import com.sun.source.tree.Tree;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Created by lwd at 2021/11/12
 *
 * @Description:
 */
public class LeetCode_103_二叉树的锯齿形层序遍历 {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> ans = new ArrayList<>();
        if (root == null) return ans;

        LinkedList<TreeNode> list = new LinkedList<>();
        list.addLast(root);

        boolean reverse = true;
        while (!list.isEmpty()) {
            ArrayList<Integer> layerAns = new ArrayList<>();
            LinkedList<TreeNode> temp = new LinkedList<>();
            while (!list.isEmpty()) {
                TreeNode node = list.pollFirst();
                if (reverse) {
                    if (node.left != null) temp.addFirst(node.left);
                    if (node.right != null) temp.addFirst(node.right);
                } else {
                    if (node.right != null) temp.addFirst(node.right);
                    if (node.left != null) temp.addFirst(node.left);
                }
                layerAns.add(node.val);
            }
            ans.add(layerAns);
            reverse = !reverse;
            list = temp;
        }
        return ans;
    }

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }

    }
}
