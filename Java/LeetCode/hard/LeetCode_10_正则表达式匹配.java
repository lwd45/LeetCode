package hard;

/**
 * Created by lwd at 2021/8/28
 *
 * @Description:
 */
public class LeetCode_10_正则表达式匹配 {
    public boolean isMatch(String s, String p) {
        int len1 = s.length(), len2 = p.length();
        boolean[][] dp = new boolean[len1 + 1][len2 + 1];
        dp[0][0] = true;
        for (int i = 2; i <= len2; i++)
            if (p.charAt(i - 1) == '*') dp[0][i] = dp[0][i - 2];

        for (int i = 1; i <= len1; i++) {
            for (int j = 1; j <= len2; j++) {
                if (p.charAt(j - 1) == '*') {
                    dp[i][j] = dp[i][j - 2] || (dp[i - 1][j - 1] && (s.charAt(i - 1) == p.charAt(j - 2) || p.charAt(j - 2) == '.'));
                } else if (p.charAt(j - 1) == '.') {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = dp[i - 1][j - 1] && s.charAt(i - 1) == p.charAt(j - 1);
                }
            }
        }
        return dp[len1][len2];
    }
}
