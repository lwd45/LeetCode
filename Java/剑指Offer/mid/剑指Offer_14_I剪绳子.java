package mid;

/**
 * Created by lwd at 2021/7/7
 *
 * @Description:
 */
public class 剑指Offer_14_I剪绳子 {
    public int cuttingRope(int n) {
        if (n <= 3) return n - 1;

        int sum = 1;
        while (n >= 3) {
            sum *= 3;
            n -= 3;
        }

        if (n > 0) {
            if (n == 1) {
                sum = sum / 3 * 4;
            } else {
                sum *= 2;
            }
        }
        return sum;
    }
}
