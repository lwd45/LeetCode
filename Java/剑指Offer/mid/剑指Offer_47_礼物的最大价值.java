package mid;

/**
 * Created by lwd at 2021/8/5
 *
 * @Description:
 */
public class 剑指Offer_47_礼物的最大价值 {
    public int maxValue(int[][] grid) {
        int r = grid.length, c = grid[0].length;
        int[][] dp = new int[r][c];

        dp[0][0] = grid[0][0];
        for (int i = 1; i < r; i++) {
            dp[i][0] = grid[i - 1][0] + grid[i][0];
        }
        for (int j = 1; j < c; j++) {
            dp[0][j] = grid[0][j - 1] + grid[0][j];
        }

        for (int i = 1; i < r; i++) {
            for (int j = 1; j < c; j++) {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]) + grid[i][j];
            }
        }
        return dp[r - 1][c - 1];
    }
}
