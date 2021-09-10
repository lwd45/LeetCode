package hard;

import java.util.LinkedList;

/**
 * Created by lwd at 2021/9/10
 *
 * @Description:
 */
public class LeetCode_239_滑动窗口最大值 {
    public int[] maxSlidingWindow(int[] nums, int k) {
        LinkedList<Integer> list = new LinkedList<Integer>();
        for (int i = 0; i < k; i++) {
            while (!list.isEmpty() && nums[i] > list.peekLast()) {
                list.pollLast();
            }
            list.addLast(nums[i]);
        }

        int[] ans = new int[nums.length - k + 1];
        ans[0] = list.peekFirst();

        for (int i = k; i < nums.length; i++) {
            if (nums[i - k] == list.peekFirst()) list.pollFirst();
            while (!list.isEmpty() && nums[i] > list.peekLast()) {
                list.pollLast();
            }
            list.addLast(nums[i]);
            ans[i - k + 1] = list.peekFirst();
        }
        return ans;
    }
}
