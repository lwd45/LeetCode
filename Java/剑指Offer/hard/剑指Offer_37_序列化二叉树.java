package hard;

import javax.swing.tree.TreeNode;
import java.util.ArrayList;
import java.util.LinkedList;

/**
 * Created by lwd at 2021/7/23
 *
 * @Description:
 */
public class 剑指Offer_37_序列化二叉树 {

    class Codec {
        public String serialize(TreeNode root) {
            LinkedList<TreeNode> list = new LinkedList<>();
            list.add(root);

            StringBuilder builder = new StringBuilder();
            while (!list.isEmpty()) {
                TreeNode poll = list.pollFirst();
                if (poll == null) {
                    builder.append("null@");
                } else {
                    builder.append(poll.val).append("@");
                    list.addLast(poll.left);
                    list.addLast(poll.right);
                }
            }
            return builder.deleteCharAt(builder.length() - 1).toString();
        }

        public TreeNode deserialize(String data) {
            String[] split = data.split("@");
            if ("null".equals(split[0])) return null;

            TreeNode treeNode = new TreeNode(Integer.parseInt(split[0]));
            LinkedList<TreeNode> list = new LinkedList<>();
            list.addLast(treeNode);
            for (int i = 1; i < split.length; ) {
                String left = split[i++];
                String right = split[i++];

                TreeNode poll = list.pollFirst();
                if (!"null".equals(left)) {
                    TreeNode lNode = new TreeNode(Integer.parseInt(left));
                    poll.left = lNode;
                    list.addLast(lNode);
                }
                if (!"null".equals(right)) {
                    TreeNode rNode = new TreeNode(Integer.parseInt(right));
                    poll.right = rNode;
                    list.addLast(rNode);
                }
            }
            return treeNode;
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
