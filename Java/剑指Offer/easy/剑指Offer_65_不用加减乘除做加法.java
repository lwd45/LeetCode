package easy;

/**
 * Created by lwd at 2021/8/6
 *
 * @Description:
 */
public class 剑指Offer_65_不用加减乘除做加法 {
    public int add(int a, int b) {
        while (b != 0) {
            int c = (a & b) << 1;
            b = a ^ b;
            a = b + c;
            b = c;
        }
        return a;
    }
}
