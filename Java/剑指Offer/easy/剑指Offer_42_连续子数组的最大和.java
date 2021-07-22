package easy;

/**
 * Created by lwd at 2021/7/22
 *
 * @Description:
 */
public class 剑指Offer_42_连续子数组的最大和 {
    public int maxSubArray(int[] nums) {
        int sum = nums[0], max = nums[0];
        for (int i = 1; i < nums.length; ++i) {
            if (sum < 0) {
                sum = nums[i];
            } else {
                sum += nums[i];
            }
            max = Math.max(max, sum);
        }
        return max;
    }
}
