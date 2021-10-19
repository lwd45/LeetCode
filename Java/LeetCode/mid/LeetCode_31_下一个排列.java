package mid;

import java.util.ArrayList;
import java.util.Arrays;

/**
 * Created by lwd at 2021/10/18
 *
 * @Description:
 */
public class LeetCode_31_下一个排列 {
    // public static void nextPermutation(int[] nums) {
    //     boolean swap = false;
    //     for (int i = nums.length - 2; i >= 0 && !swap; i--) {
    //         if (nums[i + 1] > nums[i]) {
    //             swap = true;
    //             for (int j = nums.length - 1; j > i; j--) {
    //                 if (nums[j] > nums[i]) {
    //                     int temp = nums[i];
    //                     nums[i] = nums[j];
    //                     nums[j] = temp;
    //                     Arrays.sort(nums, i + 1, nums.length);
    //                     break;
    //                 }
    //             }
    //         }
    //     }
    //     if (!swap) Arrays.sort(nums);
    // }

    public static void nextPermutation(int[] nums) {
        int index = nums.length - 2;
        while (index >= 0) {
            if (nums[index] < nums[index + 1]) {
                int r = nums.length - 1;
                while (nums[r] < nums[index]) r--;
                int temp = nums[index];
                nums[index] = nums[r];
                nums[r] = temp;
                break;
            }
            index--;
        }
        Arrays.sort(nums, index + 1, nums.length);
    }

    public static void main(String[] args) {
        nextPermutation(new int[]{1, 5, 1});
    }
}
