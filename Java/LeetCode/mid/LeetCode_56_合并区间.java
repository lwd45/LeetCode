package mid;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;

/**
 * Created by lwd at 2021/6/30
 *
 * @Description:
 */
public class LeetCode_56_合并区间 {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (o1, o2) -> o1[0] == o2[0] ? o2[1] - o1[1] : o1[0] - o2[0]);

        int start = intervals[0][0];
        int end = intervals[0][1];
        ArrayList<int[]> list = new ArrayList<>();
        for (int[] interval : intervals) {
            if (interval[0] > end) {
                list.add(new int[]{start, end});
                start = interval[0];
                end = interval[1];
            } else if (interval[1] > end) {
                end = interval[1];
            }
        }
        list.add(new int[]{start, end});

        int[][] ans = new int[list.size()][2];
        for (int i = 0; i < list.size(); ++i) {
            ans[i][0] = list.get(i)[0];
            ans[i][0] = list.get(i)[1];
        }
        return ans;
    }
}
