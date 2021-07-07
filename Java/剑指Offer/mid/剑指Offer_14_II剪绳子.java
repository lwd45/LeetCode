package mid;

/**
 * Created by lwd at 2021/7/7
 *
 * @Description:
 */
public class 剑指Offer_14_II剪绳子 {
    public static int cuttingRope(int n) {
        if (n <= 3) return n - 1;
        int mod = 1000000007;

        int a = n / 3, b = n % 3;
        long sum = 1;
        while (a > 1) {
            sum *= 3 % mod;
            a--;
        }
        if (b == 1) sum *= 4 % mod;
        else if (b == 2) sum *= 6 % mod;
        else sum *= 3 % mod;
        return (int) sum;
    }
}
