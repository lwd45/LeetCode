package mid;

/**
 * Created by lwd at 2021/8/2
 *
 * @Description:
 */
public class 剑指Offer_56_I数组中数字出现的次数 {
    public int[] singleNumbers(int[] nums) {
        int ans = 0;
        for (int num : nums) {
            ans ^= num;
        }

        int dev = 1;
        while ((dev & ans) == 0) {
            dev <<= 1;
        }

        int a = 0, b = 0;
        for (int num : nums) {
            if ((dev & num) != 0) {
                a ^= num;
            } else {
                b ^= num;
            }
        }

        return new int[]{a, b};
    }
}
