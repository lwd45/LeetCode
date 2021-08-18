package mid;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by lwd at 2021/8/18
 *
 * @Description:
 */
public class LeetCode_22_括号生成 {
    public List<String> generateParenthesis(int n) {
        ArrayList<String> ans = new ArrayList<>();
        StringBuilder builder = new StringBuilder();
        dfs(0, 0, n, ans, builder);
        return ans;
    }

    private void dfs(int left, int right, int total, ArrayList<String> ans, StringBuilder builder) {
        if (right > left || left > total) {
            return;
        }
        if (left == total && right == total) {
            String s = builder.toString();
            ans.add(s);
            return;
        }

        builder.append('(');
        dfs(left + 1, right, total, ans, builder);
        builder.deleteCharAt(builder.length() - 1);
        builder.append(')');
        dfs(left, right + 1, total, ans, builder);
        builder.deleteCharAt(builder.length() - 1);
    }
}
