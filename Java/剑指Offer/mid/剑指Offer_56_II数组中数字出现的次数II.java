package mid;

/**
 * Created by lwd at 2021/8/3
 *
 * @Description:
 */
public class 剑指Offer_56_II数组中数字出现的次数II {
    public int singleNumber(int[] nums) {
        int[] bit = new int[32];
        for (int i = 0; i < nums.length; ++i) {
            for (int j = 0; j < bit.length; ++j) {
                bit[j] += nums[i] & 1;
                nums[i] >>>= 1;
            }
        }
        int mod = 3, res = 1;
        int ans = 0;
        for (int i = 0; i < bit.length; ++i) {
            bit[i] %= mod;
            ans += res * bit[i];
            res <<= 1;
        }
        return ans;
    }
}
