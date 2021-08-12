package mid;

/**
 * Created by lwd at 2021/8/12
 *
 * @Description:
 */
public class LeetCode_238_除自身以外数组的乘积 {
    public int[] productExceptSelf(int[] nums) {
        int len = nums.length;
        int[] ans = new int[len];
        for (int i = 1; i < len; i++) {
            ans[i] = ans[i - 1] * nums[i - 1];
        }
        int right = 1;
        for (int i = len - 2; i >= 0; i--) {
            right *= nums[i + 1];
            ans[i] *= right;
        }
        return ans;
    }
}
