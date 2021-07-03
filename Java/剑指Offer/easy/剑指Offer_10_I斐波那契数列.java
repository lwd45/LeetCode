package easy;

/**
 * Created by lwd at 2021/7/3
 *
 * @Description:
 */
public class 剑指Offer_10_I斐波那契数列 {
    // public int fib(int n) {
    //     if (n == 0) return 0;
    //     else if (n == 1) return 1;
    //     return fib(n - 1) % 1000000007 + fib(n - 2) % 1000000007;
    // }

    public static int fib(int n) {
        int a = 0, b = 1, sum = 0;
        for (int i = 1; i <= n; i++) {
            sum = (a + b) % 1000000007;
            a = b;
            b = sum;
        }
        return a;
    }

}
