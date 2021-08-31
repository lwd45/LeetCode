package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/8/24
 *
 * @Description:
 */
public class LeetCode_787_K站中转内最便宜的航班 {

    public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
        int INF = 10000 * 101;
        int[][] dp = new int[k + 2][n];
        for (int i = 0; i < k + 2; i++) {
            Arrays.fill(dp[i], INF);
        }
        dp[0][src] = 0;

        for (int i = 1; i <= k + 1; i++) {
            for (int[] flight : flights) {
                int s = flight[0], d = flight[1], c = flight[2];
                dp[i][d] = Math.min(dp[i][d], dp[i - 1][s] + c);
            }
        }
        return dp[k + 1][dst] == INF ? -1 : dp[k + 1][dst];
    }

}
