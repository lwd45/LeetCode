package mid;

import javax.naming.AuthenticationNotSupportedException;

/**
 * Created by lwd at 2021/11/12
 *
 * @Description:
 */
public class 剑指Offer_66_构建乘积数组 {
    public int[] constructArr(int[] a) {
        if (a.length == 0) {
            return new int[0];
        }
        int[] ans = new int[a.length];
        ans[0] = 1;
        for (int i = 1; i < a.length; i++) {
            ans[i] = ans[i - 1] * a[i - 1];
        }
        int temp = a[a.length - 1];
        for (int i = a.length - 2; i >= 0; i--) {
            ans[i] *= temp;
            temp *= a[i];
        }
        return ans;
    }
}
