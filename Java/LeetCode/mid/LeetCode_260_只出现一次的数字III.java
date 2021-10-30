package mid;

/**
 * Created by lwd at 2021/10/30
 *
 * @Description:
 */
public class LeetCode_260_只出现一次的数字III {
    public int[] singleNumber(int[] nums) {
        int mask = nums[0];
        for (int i = 1; i < nums.length; i++) mask ^= nums[i];

        mask = mask & (-mask);
        int x1 = 0, x2 = 0;
        for (int i = 0; i < nums.length; i++) {
            if ((nums[i] & mask) != 0) {
                x1 ^= nums[i];
            } else {
                x2 ^= nums[i];
            }
        }
        return new int[]{x1, x2};
    }
}
