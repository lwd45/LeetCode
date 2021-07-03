package easy;

/**
 * Created by lwd at 2021/7/3
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *  
 * 示例 1：
 * 输入：[3,2,3]
 * 输出：3
 * <p>
 * 示例 2：
 * 输入：[2,2,1,1,1,2,2]
 * 输出：2
 *  
 * 进阶：
 * 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
 *
 * @Description:
 */
public class LeetCode_169_多数元素 {
    public int majorityElement(int[] nums) {
        int ans = nums[0];
        int count = 1;

        for (int i = 1; i < nums.length; ++i) {
            if (nums[i] != ans) {
                count--;
                if (count < 0) {
                    count = 1;
                    ans = nums[i];
                }
            } else {
                count++;
            }
        }
        return ans;
    }
}
