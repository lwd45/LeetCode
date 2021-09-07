package hard;

/**
 * Created by lwd at 2021/9/4
 *
 * @Description:
 */
public class LeetCode_4_寻找两个正序数组的中位数 {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len1 = nums1.length, len2 = nums2.length;
        if (len1 == 0 && len2 == 0) return 0;

        if ((len1 + len2) % 2 == 0) {
            return (getKthElement(nums1, nums2, (len1 + len2) / 2) + getKthElement(nums1, nums2, (len1 + len2) / 2 + 1)) / 2;
        } else {
            return getKthElement(nums1, nums2, (len1 + len2) / 2 + 1);
        }
    }

    private double getKthElement(int[] nums1, int[] nums2, int k) {
        int len1 = nums1.length, len2 = nums2.length;
        int idx1 = 0, idx2 = 0;
        while (true) {
            if (idx1 == len1) {
                return nums2[idx2 + k - 1];
            } else if (idx2 == len2) {
                return nums1[idx1 + k - 1];
            }
            if (k == 1) return Math.min(nums1[idx1], nums2[idx2]);

            int half = k / 2;
            int newIdx1 = Math.min(idx1 + half - 1, len1 - 1);
            int newIdx2 = Math.min(idx2 + half - 1, len2 - 1);
            if (nums1[newIdx1] > nums2[newIdx2]) {
                k -= (newIdx2 - idx2 + 1);
                idx2 = newIdx2 + 1;
            } else {
                k -= (newIdx1 - idx1 + 1);
                idx1 = newIdx1 + 1;
            }
        }
    }
}
