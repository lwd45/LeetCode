package mid;

import com.sun.security.auth.NTNumericCredential;

/**
 * Created by lwd at 2021/9/20
 *
 * @Description:
 */
public class LeetCode_673_最长递增子序列的个数 {
    public int findNumberOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        int[] cnt = new int[nums.length];

        int max = 1;
        for (int i = 0; i < nums.length; i++) {
            dp[i] = cnt[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    if (dp[i] < dp[j] + 1) {
                        dp[i] = dp[j] + 1;
                        cnt[i] = cnt[j];
                    } else if (dp[i] == dp[j] + 1) {
                        cnt[i] += cnt[j];
                    }
                }
            }
            max = Math.max(max, dp[i]);
        }

        int ans = 0;
        for (int i = 0; i < nums.length; i++) if (dp[i] == max) ans += cnt[i];
        return ans;
    }
}
