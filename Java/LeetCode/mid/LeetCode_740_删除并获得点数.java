package mid;

/**
 * Created by lwd at 2021/8/31
 *
 * @Description:
 */
public class LeetCode_740_删除并获得点数 {
    public int deleteAndEarn(int[] nums) {
        int max = nums[0];
        for (int num : nums) {
            max = Math.max(max, num);
        }
        int[] sum = new int[max + 1];
        for (int num : nums) {
            sum[num] += num;
        }
        return solution(sum, max);
    }

    private int solution(int[] sum, int max) {
        int first = sum[0], second = sum[1];
        for (int i = 2; i <= max; i++) {
            int temp = second;
            second = Math.max(first + sum[i], second);
            first = temp;
        }
        return Math.max(first, second);
    }
}
