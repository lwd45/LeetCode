package easy;

/**
 * Created by lwd at 2021/7/13
 *
 * @Description:
 */
public class 剑指Offer_21_调整数组顺序使奇数位于偶数前面 {
    public int[] exchange(int[] nums) {
        if (nums == null || nums.length <= 1) return nums;

        int left = 0, right = nums.length - 1;
        while (left < right) {
            while ((nums[left] & 1) == 1 && (left < right)) {
                left++;
            }
            while ((nums[right] & 1) == 0 && (left < right)) {
                right--;
            }
            int temp = nums[left];
            nums[left] = nums[right];
            nums[right] = temp;
        }
        return nums;
    }
}
