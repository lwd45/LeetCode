package mid;

import javax.lang.model.util.ElementScanner6;
import java.util.function.UnaryOperator;

/**
 * Created by lwd at 2021/8/16
 *
 * @Description:
 */
public class LeetCode_33_搜索旋转排序数组 {
    public int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) return mid;
            else if (nums[mid] >= nums[left]) {
                if (nums[mid] < target) {
                    left = mid + 1;
                } else if (target >= nums[left]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                if (nums[mid] > target) {
                    right = mid - 1;
                } else if (target > nums[right]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            }
        }
        return -1;
    }
}
