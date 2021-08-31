package easy;

/**
 * Created by lwd at 2021/8/29
 *
 * @Description:
 */
public class LeetCode_1588_所有奇数长度子数组的和 {
    public int sumOddLengthSubarrays(int[] arr) {
        int len = 1, length = arr.length;
        int ans = 0;
        for (; len <= length; len += 2) {
            int start = 0, end = start + len - 1, sum = 0;
            for (int i = start; i <= end; i++) {
                sum += arr[i];
            }
            ans += sum;

            while (end < length - 1) {
                sum -= arr[start];
                start++;
                end++;
                sum += arr[end];
                ans += sum;
            }
        }
        return ans;
    }
}
