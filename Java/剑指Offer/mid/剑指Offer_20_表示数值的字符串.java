package mid;

/**
 * Created by lwd at 2021/7/10
 * 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
 * 数值（按顺序）可以分成以下几个部分：
 * 若干空格
 * 一个 小数 或者 整数
 * （可选）一个 'e' 或 'E' ，后面跟着一个 整数
 * 若干空格
 * 小数（按顺序）可以分成以下几个部分：
 * <p>
 * （可选）一个符号字符（'+' 或 '-'）
 * 下述格式之一：
 * 至少一位数字，后面跟着一个点 '.'
 * 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
 * 一个点 '.' ，后面跟着至少一位数字
 * 整数（按顺序）可以分成以下几个部分：
 * （可选）一个符号字符（'+' 或 '-'）
 * 至少一位数字
 * 部分数值列举如下：
 * ["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
 * 部分非数值列举如下：
 * ["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
 *
 * @Description:
 */
public class 剑指Offer_20_表示数值的字符串 {
    public boolean isNumber(String s) {
        if (s == null || s.length() == 0) return false;
        boolean isNum = false, isDot = false, iseOrE = false;
        char[] cs = s.trim().toCharArray();
        for (int i = 0; i < cs.length; ++i) {
            if ('0' <= cs[i] && cs[i] <= '9') isNum = true;
            else if (cs[i] == '.') {
                if (iseOrE || isDot) return false;
                isDot = true;
            } else if (cs[i] == 'e' || cs[i] == 'E') {
                if (!isNum || iseOrE) return false;
                iseOrE = true;
                isNum = false;
            } else if (cs[i] == '-' || cs[i] == '+') {
                if (i != 0 && cs[i - 1] != 'E' && cs[i - 1] != 'e') return false;
            } else return false;
        }
        return isNum;
    }
}

















