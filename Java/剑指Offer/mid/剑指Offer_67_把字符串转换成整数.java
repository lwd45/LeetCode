package mid;

/**
 * Created by lwd at 2021/8/8
 *
 * @Description:
 */
public class 剑指Offer_67_把字符串转换成整数 {
    public static int strToInt(String str) {
        if (str == null || str.length() == 0) {
            return 0;
        }

        str = str.trim();
        boolean negative = false;
        int ans = 0;
        for (int i = 0; i < str.length(); i++) {
            if (i == 0 && ('-' == str.charAt(i) || '+' == str.charAt(i))) {
                negative = '-' == str.charAt(i);
            } else if (str.charAt(i) > '9' || str.charAt(i) < '0') {
                return negative ? -ans : ans;
            } else {
                int now = str.charAt(i) - '0';
                if (negative) {
                    if (ans > -(Integer.MIN_VALUE / 10) || (ans == -(Integer.MIN_VALUE / 10) && now > -(Integer.MIN_VALUE % 10))) {
                        return Integer.MIN_VALUE;
                    } else {
                        ans = ans * 10 + now;
                    }
                } else {
                    if (ans > Integer.MAX_VALUE / 10 || (ans == Integer.MAX_VALUE / 10 && now > Integer.MAX_VALUE % 10)) {
                        return Integer.MAX_VALUE;
                    } else {
                        ans = ans * 10 + now;
                    }
                }
            }
        }
        return negative ? -ans : ans;
    }


}
