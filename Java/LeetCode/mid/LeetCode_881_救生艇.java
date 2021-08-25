package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/8/26
 *
 * @Description:
 */
public class LeetCode_881_救生艇 {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int left = 0, right = people.length - 1;
        int ans = 0;
        while (left <= right) {
            if (people[right] + people[left] > limit) {
                right--;
            } else {
                left++;
                right--;
            }
            ans++;
        }
        return ans;
    }
}
