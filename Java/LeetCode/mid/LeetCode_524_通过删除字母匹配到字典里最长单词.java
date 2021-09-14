package mid;

import java.util.*;

/**
 * Created by lwd at 2021/9/14
 *
 * @Description:
 */
public class LeetCode_524_通过删除字母匹配到字典里最长单词 {
    public static String findLongestWord(String s, List<String> dictionary) {
        dictionary.sort((o1, o2) -> {
            if (o1.length() == o2.length()) return o2.compareTo(o1);
            return o1.length() - o2.length();
        });

        for (int i = dictionary.size() - 1; i >= 0; i--) {
            String str = dictionary.get(i);
            int k = 0, j = 0;
            for (; j < str.length(); j++) {
                while (k < s.length() && s.charAt(k) != str.charAt(j)) k++;
                if (k >= s.length()) break;
                k++;
            }
            if (j == str.length()) return str;
        }
        return "";
    }
}
