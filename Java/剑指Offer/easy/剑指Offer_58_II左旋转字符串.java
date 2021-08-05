package easy;

/**
 * Created by lwd at 2021/8/5
 *
 * @Description:
 */
public class 剑指Offer_58_II左旋转字符串 {
    public String reverseLeftWords(String s, int n) {
        return s.substring(n, s.length()) + s.substring(0, n);
    }
}
