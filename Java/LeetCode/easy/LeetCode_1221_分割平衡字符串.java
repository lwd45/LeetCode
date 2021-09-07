package easy;

/**
 * Created by lwd at 2021/9/7
 *
 * @Description:
 */
public class LeetCode_1221_分割平衡字符串 {
    public int balancedStringSplit(String s) {

        int lCount = s.charAt(0) == 'L' ? 1 : -1;
        int ans = 0;

        for (int i = 1; i < s.length(); i++) {
            lCount += s.charAt(i) == 'L' ? 1 : -1;
            if (lCount == 0) ans++;
        }
        return ans;
    }

}
