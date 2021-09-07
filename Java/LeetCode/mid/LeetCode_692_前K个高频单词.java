package mid;

import java.util.*;

/**
 * Created by lwd at 2021/9/8
 *
 * @Description:
 */
public class LeetCode_692_前K个高频单词 {
    public List<String> topKFrequent(String[] words, int k) {
        HashMap<String, Integer> cnt = new HashMap<>();
        for (String word : words) {
            int c = cnt.getOrDefault(word, 0);
            cnt.put(word, c + 1);
        }

        PriorityQueue<String> queue = new PriorityQueue<>(new Comparator<String>() {
            @Override
            public int compare(String o1, String o2) {
                if (!cnt.getOrDefault(o1, 0).equals(cnt.getOrDefault(o2, 0))) {
                    return cnt.getOrDefault(o1, 0) - cnt.getOrDefault(o2, 0);
                }
                return o1.compareTo(o2);
            }
        });

        for (String key : cnt.keySet()) {
            if (queue.size() < k) {
                queue.add(key);
            } else if (cnt.getOrDefault(key, 0) > cnt.getOrDefault(queue.peek(), 0)) {
                queue.add(key);
                queue.poll();
            }
        }

        LinkedList<String> ans = new LinkedList<>();
        while (queue.size() > 0) {
            ans.addFirst(queue.poll());
        }
        return ans;
    }
}
