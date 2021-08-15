package mid;

/**
 * Created by lwd at 2021/8/15
 *
 * @Description:
 */
public class LeetCode_678_有效的括号字符串 {
    // public boolean checkValidString(String s) {
    //     return l2r(s) && r2l(s);
    // }
    //
    // private boolean r2l(String s) {
    //     int right = 0, special = 0;
    //     for (int i = s.length() - 1; i >= 0; i--) {
    //         if (s.charAt(i) == ')') {
    //             right++;
    //         } else if (s.charAt(i) == '*') {
    //             special++;
    //         } else {
    //             if (right > 0) right--;
    //             else if (special > 0) special--;
    //             else return false;
    //         }
    //     }
    //     return right == 0 || right < special;
    // }
    //
    // private boolean l2r(String s) {
    //     int left = 0, special = 0;
    //     for (int i = 0; i < s.length(); i++) {
    //         if (s.charAt(i) == '(') {
    //             left++;
    //         } else if (s.charAt(i) == '*') {
    //             special++;
    //         } else {
    //             if (left > 0) left--;
    //             else if (special > 0) special--;
    //             else return false;
    //         }
    //     }
    //     return left == 0 || left < special;
    // }

    public boolean checkValidString(String s) {
        int left = 0, right = 0;
        for (int i = 0; i < s.length(); i++) {
            left += (s.charAt(i) == ')') ? -1 : 1;
            right += (s.charAt(s.length() - 1 - i) == '(') ? -1 : 1;
            if (left < 0 || right < 0) return false;
        }
        return true;
    }
}
