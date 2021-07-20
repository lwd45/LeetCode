package easy;

import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * Created by lwd at 2021/7/21
 *
 * @Description:
 */
public class 剑指Offer_40_最小的k个数 {
    public int[] getLeastNumbers(int[] arr, int k) {
        int[] ans = new int[k];
        if (k == 0) return null;
        PriorityQueue<Integer> queue = new PriorityQueue<>((o1, o2) -> o2 - o1);
        for (int i = 0; i < k; ++i) {
            queue.add(arr[i]);
        }
        for (int i = k; i < arr.length; ++i) {
            if (arr[i] < queue.peek()) {
                queue.poll();
                queue.add(arr[i]);
            }
        }
        for (int i = 0; i < k; ++i) {
            ans[i] = queue.poll();
        }
        return ans;
    }
}
