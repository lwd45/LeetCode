package easy;

import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * Created by lwd at 2021/8/1
 *
 * @Description:
 */
public class LeetCode_1337_矩阵中战斗力最弱的K行 {
    public int[] kWeakestRows(int[][] mat, int k) {
        PriorityQueue<int[]> queue = new PriorityQueue<>((o1, o2) -> {
            if (o2[1] == o1[1]) {
                return o2[0] - o1[0];
            }
            return o2[1] - o1[1];
        });

        for (int i = 0; i < mat.length; ++i) {
            int[] one = new int[]{i, 0};
            for (int num : mat[i]) {
                if (num != 0) one[1]++;
            }
            if (queue.size() < k) {
                queue.add(one);
            } else if (one[1] < queue.peek()[1]) {
                queue.poll();
                queue.add(one);
            }
        }

        int[] ans = new int[k];
        for (int i = k - 1; i >= 0; i--) {
            ans[i] = queue.poll()[0];
        }
        return ans;
    }
}
