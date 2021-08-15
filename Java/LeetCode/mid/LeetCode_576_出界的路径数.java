package mid;

/**
 * Created by lwd at 2021/8/15
 *
 * @Description:
 */
public class LeetCode_576_出界的路径数 {
    public int findPaths(int m, int n, int maxMove, int startRow, int startColumn) {
        int[][] dp = new int[m][n];
        int ans = 0, mod = 1000000007;
        dp[startRow][startColumn] = 1;
        for (int i = 0; i < maxMove; i++) {
            int[][] temp = new int[m][n];
            for (int r = 0; r < m; r++) {
                for (int c = 0; c < n; c++) {
                    if (dp[r][c] >= 1) {
                        if (r + 1 >= m) ans = (ans + dp[r][c]) % mod;
                        else temp[r + 1][c] = (temp[r + 1][c]+dp[r][c])% mod;

                        if (r - 1 < 0) ans = (ans + dp[r][c]) % mod;
                        else temp[r - 1][c] = (temp[r - 1][c]+dp[r][c])% mod;

                        if (c + 1 >= n) ans = (ans + dp[r][c]) % mod;
                        else temp[r][c + 1] = (temp[r][c + 1] +dp[r][c])% mod;

                        if (c - 1 < 0) ans = (ans + dp[r][c]) % mod;
                        else temp[r][c - 1] = (temp[r][c - 1] + dp[r][c])% mod;
                    }
                }
            }
            dp = temp;
        }
        return ans;
    }
}
