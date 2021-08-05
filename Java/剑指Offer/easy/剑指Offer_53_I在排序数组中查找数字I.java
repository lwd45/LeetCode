package easy;

/**
 * Created by lwd at 2021/8/5
 *
 * @Description:
 */
public class 剑指Offer_53_I在排序数组中查找数字I {
    public int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1, mid = 0;
        while (left < right) {
            mid = (left + right) / 2;
            if (nums[mid] == target) break;
            else if (nums[mid] > target) right = mid - 1;
            else left = mid + 1;
        }

        if (nums[mid] != target) return 0;
        left = mid;
        right = mid;
        while (left >= 0 && nums[left] == target) left--;
        while (right < nums.length && nums[right] == target) right++;
        return right - left - 1;
    }
}
