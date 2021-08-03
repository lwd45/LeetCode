package easy;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Created by lwd at 2021/8/3
 *
 * @Description:
 */
public class 剑指Offer_57_II和为s的连续正数序列 {
    public int[][] findContinuousSequence(int target) {
        int left = 1, right = 2, max = target / 2;

        ArrayList<int[]> ansList = new ArrayList<>();
        while (left < right && left < max) {
            int sum = (left + right) * (right - left + 1) / 2;
            if (sum > target) {
                left++;
            } else if (sum < target) {
                right++;
            } else {
                int[] ans = new int[right - left + 1];
                int index = 0;
                for (int i = left; i <= right; ++i) {
                    ans[index++] = i;
                }
                ansList.add(ans);
                left++;
            }
        }
        return ansList.toArray(new int[ansList.size()][]);
    }
}
