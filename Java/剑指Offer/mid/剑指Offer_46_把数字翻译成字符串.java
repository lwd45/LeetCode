package mid;

import javax.print.attribute.standard.NumberUp;

/**
 * Created by lwd at 2021/8/4
 *
 * @Description:
 */
public class 剑指Offer_46_把数字翻译成字符串 {


    // public int translateNum(int num) {
    //     String s = String.valueOf(num);
    //     int[] nums = new int[s.length()];
    //     int index = 0;
    //     for (char c : s.toCharArray()) {
    //         nums[index++] = Integer.parseInt(String.valueOf(c));
    //     }
    //     return dfs(nums, 0);
    // }
    //
    // public int dfs(int[] nums, int index) {
    //     if (index == nums.length || index == nums.length - 1) {
    //         return 1;
    //     }
    //
    //     if (nums[index] != 0 && nums[index] <= 2) {
    //         if (nums[index] == 1 || nums[index + 1] <= 5) {
    //             return dfs(nums, index + 1) + dfs(nums, index + 2);
    //         }
    //     }
    //     return dfs(nums, index + 1);
    // }

    public int translateNum(int num) {
        if (num < 10) return 1;
        if (num % 100 < 26 && num % 100 >= 10) {
            return translateNum(num / 100) + translateNum(num / 10);
        }
        return translateNum(num / 10);
    }
}
