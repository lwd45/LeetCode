package easy;

import java.util.Arrays;

/**
 * Created by lwd at 2021/8/8
 *
 * @Description:
 */
public class 剑指Offer_61_扑克牌中的顺子 {
    public boolean isStraight(int[] nums) {
        Arrays.sort(nums);
        int index = 0, zero = 0;
        for (; index < nums.length && nums[index] == 0; index++) {
            zero++;
        }
        int i = index;
        for (; index < nums.length; index++) {
            if (index > 0 && nums[index] == nums[index - 1]) {
                return false;
            }
        }
        return (nums[nums.length - 1] - nums[i] - (nums.length - 1 - i)) <= zero;
    }
}
