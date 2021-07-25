package easy;

/**
 * Created by lwd at 2021/7/25
 *
 * @Description:
 */
public class 剑指Offer_39_数组中出现次数超过一半的数字 {

    public int majorityElement(int[] nums) {
        int count = 1, ans = nums[0];
        for (int i = 1; i < nums.length; ++i) {
            if (nums[i] == ans) {
                count++;
            } else {
                count--;
                if (count < 0) {
                    ans = nums[i];
                    count = 1;
                }
            }
        }
        return ans;
    }
}
