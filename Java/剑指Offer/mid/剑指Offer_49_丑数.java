package mid;

/**
 * Created by lwd at 2021/8/6
 *
 * @Description:
 */
public class 剑指Offer_49_丑数 {
    public int nthUglyNumber(int n) {
        int[] nums = new int[n];
        nums[0] = 1;
        int index2 = 0, index3 = 0, index5 = 0;
        for (int i = 1; i < n; i++) {
            nums[i] = Math.min(2 * nums[index2], Math.min(3 * nums[index3], 5 * nums[index5]));
            if (nums[i] / nums[index2] == index2 && nums[i] % nums[index2] == 0) index2++;
            if (nums[i] / nums[index3] == index3 && nums[i] % nums[index3] == 0) index3++;
            if (nums[i] / nums[index5] == index5 && nums[i] % nums[index5] == 0) index5++;
        }
    }
}
