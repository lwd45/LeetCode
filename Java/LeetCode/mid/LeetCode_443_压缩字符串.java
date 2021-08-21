package mid;

/**
 * Created by lwd at 2021/8/21
 *
 * @Description:
 */
public class LeetCode_443_压缩字符串 {
    public int compress(char[] chars) {
        if (chars.length == 1) return 1;
        int index = 0, left = 0, right = 0, count = 0;
        while (right < chars.length) {
            if (chars[right] == chars[left]) {
                count++;
            } else {
                chars[index] = chars[left];
                index++;
                if (count > 1) {
                    String str = count + "";
                    for (int i = 0; i < str.length(); i++) {
                        chars[index] = str.charAt(i);
                        index++;
                    }
                }
                left = right;
                count = 1;
            }
            right++;
        }

        chars[index] = chars[left];
        index++;
        if (count > 1) {
            String str = count + "";
            for (int i = 0; i < str.length(); i++) {
                chars[index] = str.charAt(i);
                index++;
            }
        }
        return index;
    }
}
