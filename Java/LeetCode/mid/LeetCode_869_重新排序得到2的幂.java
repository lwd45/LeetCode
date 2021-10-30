package mid;

import java.util.HashMap;
import java.util.HashSet;

/**
 * Created by lwd at 2021/10/28
 * 10
 * 1010
 * <p>
 * 46
 * 10 1110
 *
 * @Description:
 */
public class LeetCode_869_重新排序得到2的幂 {
    public boolean reorderedPowerOf2(int n) {
        HashSet<String> set = new HashSet<>();
        int num = 1;
        while (num < 1000000000) {
            set.add(helper(num));
            num = num << 1;
        }
        return set.contains(helper(n));
    }

    public String helper(int n) {
        char[] cs = new char[10];
        while (n > 0) {
            int temp = n % 10;
            cs[temp]++;
            n = n / 10;
        }
        return new String(cs);
    }

}
