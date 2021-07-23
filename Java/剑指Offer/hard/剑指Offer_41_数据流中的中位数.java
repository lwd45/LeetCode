package hard;

import java.util.PriorityQueue;

/**
 * Created by lwd at 2021/7/23
 *
 * @Description:
 */
public class 剑指Offer_41_数据流中的中位数 {

    class MedianFinder {

        private final PriorityQueue<Integer> minQueue;//奇数多一位 存后半部分
        private final PriorityQueue<Integer> maxQueue;//存前半部分

        public MedianFinder() {
            minQueue = new PriorityQueue<>();
            maxQueue = new PriorityQueue<>((x, y) -> (y - x));
        }

        public void addNum(int num) {
            if (minQueue.size() == maxQueue.size()) {
                maxQueue.add(num);
                minQueue.add(maxQueue.poll());
            } else {
                minQueue.add(num);
                maxQueue.add(minQueue.poll());
            }
        }

        public double findMedian() {
            return minQueue.size() == maxQueue.size() ? (double) (minQueue.peek() + maxQueue.peek()) / 2 : minQueue.peek();
        }
    }
}
