package easy;

/**
 * Created by lwd at 2021/8/20
 *
 * @Description:
 */
public class LeetCode_541_反转字符串II {
    public String reverseStr(String s, int k) {
        int length = s.length();

        char[] cs = s.toCharArray();
        for (int i = 0; i < length; ) {
            int start, end;
            if (i + 2 * k <= length) {
                start = i;
                end = i + k - 1;
            } else if (i + k <= length) {
                start = i;
                end = i + k - 1;
            } else {
                start = i;
                end = length - 1;
            }
            reverse(start, end, cs);
            i += 2 * k;
        }
        return new String(cs);
    }

    private void reverse(int start, int end, char[] cs) {
        while (start < end) {
            char c = cs[start];
            cs[start] = cs[end];
            cs[end] = c;
            start++;
            end--;
        }
    }
}
