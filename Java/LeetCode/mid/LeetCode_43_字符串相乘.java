package mid;

/**
 * Created by lwd at 2021/8/20
 *
 * @Description:
 */
public class LeetCode_43_字符串相乘 {
    public static String multiply(String num1, String num2) {
        if (num1 == null || num2 == null || num1.length() == 0 || num2.length() == 0 || "0".equals(num1) || "0".equals(num2))
            return "0";

        int len2 = num2.length();
        String ans = "0";
        for (int i = len2 - 1; i >= 0; i--) {
            String multi = multi(num1, num2.charAt(i) - '0', len2 - 1 - i);
            ans = add(ans, multi);
        }
        return ans;
    }

    private static String multi(String num1, int c, int zero) {
        StringBuilder builder = new StringBuilder();

        int plus = 0;
        for (int i = num1.length() - 1; i >= 0; i--) {
            int c1 = num1.charAt(i) - '0';
            int ans = (c * c1 + plus) % 10;
            plus = (c * c1 + plus) / 10;
            builder.append(ans);
        }
        if (plus > 0) builder.append(plus);
        builder.reverse();
        while (zero-- > 0) {
            builder.append('0');
        }
        return builder.toString();
    }

    private static String add(String ans, String multi) {
        StringBuilder builder = new StringBuilder();

        int len1 = ans.length() - 1, len2 = multi.length() - 1;
        int plus = 0;
        while (len1 >= 0 && len2 >= 0) {
            int i = ans.charAt(len1) - '0';
            int j = multi.charAt(len2) - '0';
            int temp = (i + j + plus) % 10;
            plus = (i + j + plus) / 10;
            builder.append(temp);
            len1--;
            len2--;
        }
        while (len1 >= 0) {
            int i = ans.charAt(len1) - '0';
            int temp = (i + plus) % 10;
            plus = (i + plus) / 10;
            builder.append(temp);
            len1--;
        }
        while (len2 >= 0) {
            int i = multi.charAt(len2) - '0';
            int temp = (i + plus) % 10;
            plus = (i + plus) / 10;
            builder.append(temp);
            len2--;
        }
        if (plus > 0) {
            builder.append(1);
        }
        return builder.reverse().toString();
    }

    public static void main(String[] args) {
        System.out.println(multiply("123", "456"));
    }
}
