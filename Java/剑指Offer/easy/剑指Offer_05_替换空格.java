package easy;

/**
 * Created by lwd at 2021/7/5
 *
 * @Description:
 */
public class 剑指Offer_05_替换空格 {
    public String replaceSpace(String s) {
        StringBuilder builder = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (c == ' ') {
                builder.append("%20");
            } else {
                builder.append(c);
            }
        }
        return builder.toString();
    }
}
