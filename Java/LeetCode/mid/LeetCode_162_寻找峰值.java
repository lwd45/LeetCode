package mid;

/**
 * Created by lwd at 2021/8/16
 *
 * @Description:
 */
public class LeetCode_162_寻找峰值 {
    // public int findPeakElement(int[] nums) {
    //     if (nums.length == 1) return 0;
    //     else if (nums[0] > nums[1]) return 0;
    //     else if (nums[nums.length - 1] > nums[nums.length - 2]) return nums.length - 1;
    //
    //     for (int i = 1; i < nums.length - 1; i++) {
    //         if (nums[i] > nums[i - 1] && nums[i] > nums[i + 1]) return i;
    //     }
    //     return -1;
    // }

    public int findPeakElement(int[] nums) {
        int left = 0, right = nums.length - 1;
        while (left < right) {
            int mid = (left + right) / 2;
            if (nums[mid] > nums[mid + 1]) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left;
    }
}
