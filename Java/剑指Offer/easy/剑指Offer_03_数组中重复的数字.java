package easy;

/**
 * Created by lwd at 2021/7/3
 * 找出数组中重复的数字。
 * 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
 *
 * <p>
 * 示例 1：
 * 输入：
 * [2, 3, 1, 0, 2, 5, 3]
 * 输出：2 或 3
 *  
 * <p>
 * 限制：
 * 2 <= n <= 100000
 *
 * @Description:
 */
public class 剑指Offer_03_数组中重复的数字 {
    // public int findRepeatNumber(int[] nums) {
    //     boolean[] visited = new boolean[nums.length];
    //     for (int num : nums) {
    //         if (visited[num]) return num;
    //         visited[num] = true;
    //     }
    //     return 0;
    // }

    public int findRepeatNumber(int[] nums) {
        int i = 0;
        while (i < nums.length) {
            if (nums[i] == i) {
                i++;
                continue;
            }

            if (nums[nums[i]] == nums[i]) {
                return nums[i];
            } else {
                int temp = nums[i];
                nums[i] = nums[temp];
                nums[temp] = temp;
            }
        }
        return -1;
    }
}
