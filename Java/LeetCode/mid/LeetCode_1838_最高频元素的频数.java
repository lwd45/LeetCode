package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/7/19
 *
 * @Description:
 */
public class LeetCode_1838_最高频元素的频数 {
    public int maxFrequency(int[] nums, int k) {
        int left = 0, right = 1, ans = 1, sum = 0;
        Arrays.sort(nums);
        for (; right < nums.length; ++right) {
            int add = (right - left) * (nums[right] - nums[right - 1]);
            sum += add;
            if (sum > k) {
                while (sum > k && left < right) {
                    sum -= nums[right] - nums[left];
                    left++;
                }
            } else {
                ans = Math.max(right - left + 1, ans);
            }
        }
        return ans;
    }
}
