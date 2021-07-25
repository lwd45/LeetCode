package mid;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/**
 * Created by lwd at 2021/7/25
 * [4,-2],[1,4],[-3,1]
 * <p>
 * [-3,1]
 * [-2,4]
 * [ 1,4]
 *
 * @Description:
 */
public class LeetCode_1743_从相邻元素对还原数组 {
    public int[] restoreArray(int[][] adjacentPairs) {
        HashMap<Integer, List<Integer>> map = new HashMap<>();
        ArrayList<Integer> startList = new ArrayList<>();
        for (int[] adjacentPair : adjacentPairs) {
            int i = adjacentPair[0];
            int j = adjacentPair[1];
            List<Integer> listI = map.getOrDefault(i, new ArrayList<>());
            listI.add(j);
            map.put(i, listI);

            List<Integer> listJ = map.getOrDefault(j, new ArrayList<>());
            listJ.add(i);
            map.put(j, listJ);
        }

        Integer start = 0;
        for (Integer key : map.keySet()) {
            if (map.get(key).size() == 1) {
                start = key;
                break;
            }
        }

        int[] ans = new int[adjacentPairs.length + 1];
        ans[0] = start;
        for (int i = 1; i < ans.length; ++i) {
            List<Integer> list = map.get(ans[i - 1]);
            Integer next = list.get(0);
            list.remove((Integer) next);
            map.put(ans[i - 1], list);
            ans[i] = next;
            List<Integer> list1 = map.get(next);
            list1.remove((Integer) ans[i - 1]);
            map.put(next, list1);
        }
        return ans;
    }
}
