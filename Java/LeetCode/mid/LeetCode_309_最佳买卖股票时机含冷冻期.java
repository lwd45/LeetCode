package mid;

/**
 * Created by lwd at 2021/8/22
 *
 * @Description:
 */
public class LeetCode_309_最佳买卖股票时机含冷冻期 {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length < 2) return 0;

        int[][] dp = new int[prices.length][3];
        dp[0][0] = -prices[0];//持有
        dp[0][1] = 0;//不持有，但是不冷冻
        dp[0][2] = 0;//不持有，冷冻
        for (int i = 1; i < prices.length; i++) {
            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] - prices[i]);
            dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][2]);
            dp[i][2] = dp[i - 1][0] + prices[i];
        }
        return Math.max(dp[prices.length - 1][1], dp[prices.length - 1][2]);
    }
}
