package mid;

import java.util.HashSet;

/**
 * Created by lwd at 2021/9/3
 *
 * @Description:
 */
public class LeetCode_128_最长连续序列 {
    public int longestConsecutive(int[] nums) {
        HashSet<Integer> set = new HashSet<Integer>();
        for (int num : nums) set.add(num);
        int ans = 0;
        for (int num : nums) {
            if (set.contains(num - 1)) {
                continue;
            } else {
                int count = 1, plus = 1;
                while (set.contains(num + plus)) {
                    count++;
                    plus++;
                }
                ans = Math.max(ans, count);
            }
        }
        return ans;
    }
}
