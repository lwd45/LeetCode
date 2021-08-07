package hard;

import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * Created by lwd at 2021/8/6
 *
 * @Description:
 */
public class 剑指Offer_59_I滑动窗口的最大值 {
    public int[] maxSlidingWindow(int[] nums, int k) {
        PriorityQueue<int[]> queue = new PriorityQueue<>((int[] num1, int[] num2) -> {
            return num1[0] != num2[0] ? num2[0] - num1[0] : num2[1] - num1[1];
        });

        int[] ans = new int[nums.length - k + 1];
        for (int i = 0; i < k; ++i) {
            queue.add(new int[]{nums[i], i});
        }
        ans[0] = queue.peek()[0];

        for (int i = k; i < nums.length; i++) {
            queue.add(new int[]{nums[i], i});
            while (queue.peek()[1] <= i - k) {
                queue.poll();
            }
            ans[i - k + 1] = queue.peek()[0];
        }
        return ans;
    }
}
