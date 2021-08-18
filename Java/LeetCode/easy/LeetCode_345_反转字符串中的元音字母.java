package easy;

import java.util.ArrayList;
import java.util.Arrays;

/**
 * Created by lwd at 2021/8/19
 *
 * @Description:
 */
public class LeetCode_345_反转字符串中的元音字母 {
    public String reverseVowels(String s) {
        char[] chars = s.toCharArray();
        ArrayList<Character> characters = new ArrayList<>();
        characters.add('a');
        characters.add('o');
        characters.add('e');
        characters.add('i');
        characters.add('u');
        characters.add('A');
        characters.add('O');
        characters.add('E');
        characters.add('I');
        characters.add('U');

        int left = 0, right = chars.length - 1;
        while (left < right) {
            while (left < right && !characters.contains(chars[left])) left++;
            while (left < right && !characters.contains(chars[right])) right--;
            char temp = chars[left];
            chars[left] = chars[right];
            chars[right] = temp;
        }
        return new String(chars);
    }
}
