package mid;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Created by lwd at 2021/8/18
 *
 * @Description:
 */
public class LeetCode_199_二叉树的右视图 {
    public List<Integer> rightSideView(TreeNode root) {
        ArrayList<Integer> ans = new ArrayList<>();
        if (root == null) return ans;
        LinkedList<TreeNode> layer = new LinkedList<>();
        layer.addLast(root);
        while (!layer.isEmpty()) {
            int size = layer.size();
            ans.add(layer.get(size - 1).val);
            while (size > 0) {
                TreeNode first = layer.removeFirst();
                if (first.left != null) layer.addLast(first.left);
                if (first.right != null) layer.addLast(first.right);
                size--;
            }
        }
        return ans;
    }

    class TreeNode {
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
