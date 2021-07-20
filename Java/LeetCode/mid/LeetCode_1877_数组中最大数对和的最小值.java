package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/7/20
 *
 * @Description:
 */
public class LeetCode_1877_数组中最大数对和的最小值 {
    public int minPairSum(int[] nums) {
        Arrays.sort(nums);
        int ans = 0;
        for (int i = 0, j = nums.length - 1; i < j; ++i, --j) {
            ans = Math.max(ans, nums[i] + nums[j]);
        }
        return ans;
    }
}
