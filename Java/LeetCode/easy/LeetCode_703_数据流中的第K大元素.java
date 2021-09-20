package easy;

import java.util.PriorityQueue;

/**
 * Created by lwd at 2021/9/18
 *
 * @Description:
 */
public class LeetCode_703_数据流中的第K大元素 {
    class KthLargest {
        PriorityQueue<Integer> minHeap;
        Integer k;

        public KthLargest(int k, int[] nums) {
            this.k = k;
            minHeap = new PriorityQueue<>();
            for (int num : nums) {
                if (minHeap.size() < k) {
                    minHeap.add(num);
                } else if (num > minHeap.peek()) {
                    minHeap.poll();
                    minHeap.add(num);
                }
            }
        }

        public int add(int val) {
            if (minHeap.size() < k) {
                minHeap.add(val);
            } else if (val > minHeap.peek()) {
                minHeap.poll();
                minHeap.add(val);
            }
            return minHeap.peek();
        }
    }
}
