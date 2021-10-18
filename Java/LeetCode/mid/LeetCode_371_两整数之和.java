package mid;

/**
 * Created by lwd at 2021/9/26
 *
 * @Description:
 */
public class LeetCode_371_两整数之和 {
    public int getSum(int a, int b) {
        while ((a & b) != 0) {
            int x = a & b;
            int y = a ^ b;
            a = x << 1;
            b = y;
        }
        return a ^ b;
    }
}
