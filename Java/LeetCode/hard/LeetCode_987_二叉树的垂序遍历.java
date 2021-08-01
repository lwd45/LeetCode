package hard;

import com.sun.org.apache.bcel.internal.generic.NEW;

import java.util.*;

/**
 * Created by lwd at 2021/7/31
 *
 * @Description:
 */
public class LeetCode_987_二叉树的垂序遍历 {

    // public List<List<Integer>> verticalTraversal(TreeNode root) {
    //     HashMap<String, List<Integer>> map1 = new HashMap<>();//row@col, list
    //     dfs(root, 0, 0, map1);
    //
    //     Set<String> temp = map1.keySet();
    //     ArrayList<String> keyList = new ArrayList<>(temp);
    //     keyList.sort((o1, o2) -> {
    //         int col1 = Integer.parseInt(o1.split("@")[0]);
    //         int row1 = Integer.parseInt(o1.split("@")[1]);
    //         int col2 = Integer.parseInt(o2.split("@")[0]);
    //         int row2 = Integer.parseInt(o2.split("@")[1]);
    //         if (col1 == col2) {
    //             return row1 - row2;
    //         }
    //         return col1 - col2;
    //     });
    //
    //     HashMap<Integer, List<Integer>> map = new HashMap<>();//col -- list)
    //     for (String key : keyList) {
    //         Integer k = Integer.parseInt(key.split("@")[0]);
    //         List<Integer> integers = map.getOrDefault(k, new ArrayList<>());
    //         List<Integer> list = map1.get(key);
    //         Collections.sort(list);
    //         integers.addAll(list);
    //         map.put(k, integers);
    //     }
    //
    //     ArrayList<List<Integer>> ans = new ArrayList<>();
    //     HashSet<Integer> set = new HashSet<>();
    //     for (String key : keyList) {
    //         Integer k = Integer.parseInt(key.split("@")[0]);
    //         if (!set.contains(k)) {
    //             set.add(k);
    //             ans.add(map.get(k));
    //         }
    //     }
    //     return ans;
    // }

    // private void dfs(TreeNode root, int col, int row, HashMap<String, List<Integer>> map1) {
    //     if (root == null) return;
    //     String key = col + "@" + row;
    //     List<Integer> list = map1.getOrDefault(key, new ArrayList<>());
    //     list.add(root.val);
    //     map1.put(key, list);
    //
    //     dfs(root.left, col - 1, row + 1, map1);
    //     dfs(root.right, col + 1, row + 1, map1);
    // }


    public List<List<Integer>> verticalTraversal(TreeNode root) {
        List<int[]> temp = new ArrayList<>();
        dfs(root, 0, 0, temp);
        temp.sort((int[] list1, int[] list2) -> {
            if (list1[0] != list2[0]) {
                return list1[0] - list2[0];
            } else if (list1[1] != list2[1]) {
                return list1[1] - list2[1];
            }
            return list1[2] - list2[2];
        });

        List<List<Integer>> ans = new ArrayList<>();
        int index = 0;
        int lastIndex = Integer.MIN_VALUE;
        for (int[] t : temp) {
            if (t[0] != lastIndex) {
                lastIndex = t[0];
                ans.add(new ArrayList<>());
                index++;
            }
            ans.get(index - 1).add(t[2]);
        }
        return ans;
    }

    public void dfs(TreeNode root, int row, int col, List<int[]> temp) {
        if (root == null) return;
        temp.add(new int[]{col, row, root.val});
        dfs(root.left, row + 1, col - 1, temp);
        dfs(root.right, row + 1, col + 1, temp);
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
