package hard;

/**
 * Created by lwd at 2021/11/2
 *
 * @Description:
 */
public class LeetCode_115_不同的子序列 {
    public int numDistinct(String s, String t) {
        int len1 = s.length(), len2 = t.length();
        if (len1 < len2) return 0;
        int[][] dp = new int[len1 + 1][len2 + 1];
        for (int i = 0; i <= len1; i++) dp[i][0] = 1;

        for (int i = 1; i <= len1; i++) {
            for (int j = 1; j <= len2; j++) {
                if (s.charAt(i - 1) == t.charAt(j - 1)) dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
                else dp[i][j] = dp[i - 1][j];
            }
        }
        return dp[len1][len2];
    }
}
