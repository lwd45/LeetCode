package mid;

import java.util.HashMap;
import java.util.List;
import java.util.Set;

/**
 * Created by lwd at 2021/9/13
 *
 * @Description:
 */
public class LeetCode_447_回旋镖的数量 {
    public int numberOfBoomerangs(int[][] points) {

        for (int[] p : points) {
            HashMap<Integer, Integer> cnt = new HashMap<>();

            for (int[] q : points) {
                int dis = (p[0] - q[0]) * (p[0] - q[0]) + (p[1] - q[1]) * (p[1] - q[1]);
                cnt.put(dis, cnt.getOrDefault(dis, 0) + 1);
            }
        }

        int ans = 0;
        for (Integer key : cnt.keySet()) {
            int m = cnt.get(key);
            ans += m * (m - 1);
        }
        return ans;
    }
}

