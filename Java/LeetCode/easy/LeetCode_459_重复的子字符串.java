package easy;

/**
 * Created by lwd at 2021/8/29
 * abc abc abc
 *
 * @Description:
 */
public class LeetCode_459_重复的子字符串 {
    public boolean repeatedSubstringPattern(String s) {
        if (s.length() == 1) return true;
        int len = 1, length = s.length();
        while (len <= length / 2) {
            for (int i = 0; i < length; i++) {
                if (s.charAt(i) == s.charAt(i % len) && i == length - 1 && (i % len) == len - 1) return true;
                else if (s.charAt(i) != s.charAt(i % len)) break;
            }
            len++;
        }
        return false;
    }
}
