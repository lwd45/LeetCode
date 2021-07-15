package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/7/15
 *
 * @Description:
 */
public class LeetCode_1846_减小和重新排列数组后的最大元素 {
    // public int maximumElementAfterDecrementingAndRearranging(int[] arr) {
    //     Arrays.sort(arr);
    //     arr[0] = 1;
    //     for (int i = 1; i < arr.length; ++i) {
    //         arr[i] = Math.min(arr[i - 1] + 1, arr[i]);
    //     }
    //     return arr[arr.length - 1];
    // }

    public int maximumElementAfterDecrementingAndRearranging(int[] arr) {
        int[] count = new int[arr.length + 1];
        int len = arr.length;
        for (int v : arr) {
            ++count[Math.min(v, len)];
        }

        int miss = 0;
        for (int i = 1; i < count.length; ++i) {
            if (count[i] == 0) {
                miss++;
            } else {
                miss -= Math.min((count[i] - 1), miss);
            }
        }
        return len - miss;
    }
}
