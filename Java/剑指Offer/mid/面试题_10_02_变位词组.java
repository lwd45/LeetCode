package mid;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/**
 * Created by lwd at 2021/7/18
 *
 * @Description:
 */
public class 面试题_10_02_变位词组 {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String>> map = new HashMap<>();

        for (String str : strs) {
            int[] cnt = new int[26];
            char[] chars = str.toCharArray();
            for (char c : chars) {
                cnt[c - 'a']++;
            }

            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < 26; ++i) {
                if (cnt[i] != 0) {
                    sb.append('a' + i);
                    sb.append(cnt[i]);
                }
            }
            List<String> list = map.getOrDefault(sb.toString(), new ArrayList<String>());
            list.add(str);
            map.put(sb.toString(), list);
        }

        ArrayList<List<String>> ans = new ArrayList<>();
        for (String key : map.keySet()) {
            ans.add(map.get(key));
        }
        return ans;
    }
}
