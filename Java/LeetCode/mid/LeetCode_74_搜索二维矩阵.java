package mid;

/**
 * Created by lwd at 2021/8/13
 *
 * @Description:
 */
public class LeetCode_74_搜索二维矩阵 {
    public boolean searchMatrix(int[][] matrix, int target) {
        int row = matrix.length, col = matrix[0].length;
        int i = row - 1, j = 0;
        while (i >= 0 && j < col) {
            if (matrix[i][j] == target) return true;
            else if (matrix[i][j] > target) i--;
            else j++;
        }
        return false;
    }
}
