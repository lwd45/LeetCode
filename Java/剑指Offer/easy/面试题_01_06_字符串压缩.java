package easy;

/**
 * Created by lwd at 2021/8/12
 *
 * @Description:
 */
public class 面试题_01_06_字符串压缩 {

    public String compressString(String S) {
        StringBuilder sb = new StringBuilder();

        int count = 1;
        for (int i = 1; i < S.length(); i++) {
            if (S.charAt(i) == S.charAt(i - 1)) {
                count++;
            } else {
                sb.append(S.charAt(i - 1));
                sb.append(count);
                count = 1;
            }
        }
        sb.append(S.charAt(S.length() - 1));
        sb.append(count);

        if (sb.length() > S.length()) {
            return S;
        }
        return sb.toString();
    }

}
