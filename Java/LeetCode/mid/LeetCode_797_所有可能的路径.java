package mid;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Created by lwd at 2021/8/25
 *
 * @Description:
 */
public class LeetCode_797_所有可能的路径 {
    public List<List<Integer>> allPathsSourceTarget(int[][] graph) {
        ArrayList<List<Integer>> ans = new ArrayList<>();
        LinkedList<Integer> list = new LinkedList<>();
        list.addLast(0);
        dfs(ans, list, 0, graph, graph.length - 1);
        return ans;
    }

    private void dfs(ArrayList<List<Integer>> ans, LinkedList<Integer> list, int nowNum, int[][] graph, int target) {
        if (nowNum == target) {
            ans.add(new ArrayList<>(list));
            return;
        }

        for (int g : graph[nowNum]) {
            list.addLast(g);
            dfs(ans, list, g, graph, target);
            list.removeLast();
        }
    }
}
