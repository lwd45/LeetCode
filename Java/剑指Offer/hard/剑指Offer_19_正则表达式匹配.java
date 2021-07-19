package hard;

/**
 * Created by lwd at 2021/7/19
 *
 * @Description:
 */
public class 剑指Offer_19_正则表达式匹配 {
    public boolean isMatch(String s, String p) {
        int sLen = s.length(), pLen = p.length();
        boolean[][] match = new boolean[sLen + 1][pLen + 1];
        for (int i = 0; i <= sLen; ++i) {
            for (int j = 0; j <= pLen; ++j) {
                if (j == 0) {
                    match[i][j] = i == 0;
                } else {
                    if (p.charAt(j - 1) == '*') {
                        if (j >= 2 && i > 0 && (s.charAt(i - 1) == p.charAt(j - 2) || p.charAt(j - 2) == '.')) {
                            match[i][j] |= match[i - 1][j];
                        }
                        if (j >= 2) {
                            match[i][j] |= match[i][j - 2];
                        }
                    } else if (p.charAt(j - 1) == '.' && i > 0) {
                        match[i][j] = match[i - 1][j - 1];
                    } else if (i > 0) {
                        match[i][j] = match[i - 1][j - 1] && s.charAt(i - 1) == p.charAt(j - 1);
                    }
                }
            }
        }
        return match[sLen][pLen];
    }
}
