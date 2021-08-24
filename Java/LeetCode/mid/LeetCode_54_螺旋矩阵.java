package mid;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by lwd at 2021/8/24
 *
 * @Description:
 */
public class LeetCode_54_螺旋矩阵 {
    public List<Integer> spiralOrder(int[][] matrix) {
        int rowStart = 0, rowEnd = matrix.length - 1, colStart = 0, colEnd = matrix[0].length - 1;
        List<Integer> ans = new ArrayList<>();

        int dIndex = 0;
        int[][] dir = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

        int i = 0, j = 0;
        while (rowStart <= rowEnd && colStart <= colEnd) {
            while (rowStart <= i && i <= rowEnd && colStart <= j && j <= colEnd) {
                ans.add(matrix[i][j]);
                i += dir[dIndex][0];
                j += dir[dIndex][1];
            }
            i -= dir[dIndex][0];
            j -= dir[dIndex][1];

            if (dIndex == 0) {
                rowStart++;
            } else if (dIndex == 1) {
                colEnd--;
            } else if (dIndex == 2) {
                rowEnd--;
            } else {
                colStart++;
            }

            dIndex = (dIndex + 1) % 4;
            i += dir[dIndex][0];
            j += dir[dIndex][1];
        }
        return ans;
    }
}
