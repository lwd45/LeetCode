package mid;

/**
 * Created by lwd at 2021/8/10
 *
 * @Description:
 */
public class LeetCode_413_等差数列划分 {
    public int numberOfArithmeticSlices(int[] nums) {
        if (nums == null || nums.length < 3) return 0;

        int length = nums.length, ans = 0, d = nums[1] - nums[0], t = 0;
        for (int i = 2; i < length; i++) {
            if (nums[i] - nums[i - 1] == d) {
                t++;
            } else {
                d = nums[i] - nums[i - 1];
                t = 0;
            }
            ans += t;
        }
        return ans;
    }
}
