package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/7/14
 *
 * @Description:
 */
public class LeetCode_1818_绝对差值和 {
    public int minAbsoluteSumDiff(int[] nums1, int[] nums2) {
        int[] copy = new int[nums1.length];
        System.arraycopy(nums1, 0, copy, 0, nums1.length);
        Arrays.sort(copy);

        int sum = 0, mod = 1000000007, max = 0;
        for (int i = 0; i < nums1.length; ++i) {
            int gap = Math.abs(nums1[i] - nums2[i]);
            sum = (sum + gap) % mod;

            int j = binarySearch(copy, nums2[i]);
            if (j < nums1.length) {
                max = Math.max(max, gap - Math.abs(copy[j] - nums2[i]));
            }
            if (j > 0) {
                max = Math.max(max, gap - Math.abs(copy[j - 1] - nums2[i]));
            }
        }
        return (sum - max + mod) % mod;
    }

    public int binarySearch(int[] src, int target) {
        int left = 0, right = src.length - 1;
        while (left < right) {
            int mid = (left + right) >> 1;
            if (src[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return left;//nums[left]>=target 第一个满足该条件的left
    }
}
