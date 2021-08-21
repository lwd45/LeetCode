package mid;

/**
 * Created by lwd at 2021/8/21
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 * 你可以认为每种硬币的数量是无限的。
 * <p>
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * <p>
 * 输入：coins = [2], amount = 3
 * 输出：-1
 * <p>
 * 输入：coins = [1], amount = 0
 * 输出：0
 * <p>
 * 输入：coins = [1], amount = 1
 * 输出：1
 * <p>
 * 输入：coins = [1], amount = 2
 * 输出：2
 *  
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 231 - 1
 * 0 <= amount <= 104
 *
 * @Description:
 */
public class LeetCode_322_零钱兑换 {
    public int coinChange(int[] coins, int amount) {
        if (amount == 0) return 0;
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, Integer.MAX_VALUE - 1);

        for (int i = 1; i <= amount; i++) {
            for (int coin : coins) {
                if (i == coin) dp[i] = 1;
                else if (i - coin > 0 && dp[i - coin] > 0) dp[i] = Math.min(dp[i - coin] + 1, dp[i]);
            }
        }

        return dp[amount] == Integer.MAX_VALUE-1 ? -1 : dp[amount];
    }
}
