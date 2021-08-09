package easy;

/**
 * Created by lwd at 2021/8/8
 *
 * @Description:
 */
public class 剑指Offer_62_圆圈中最后剩下的数字 {
    public int lastRemaining(int n, int m) {
        int ans = 0;
        for (int i = 2; i <= n; i++) {
            ans = (ans + m) % i;
        }
        return ans;
    }
}
