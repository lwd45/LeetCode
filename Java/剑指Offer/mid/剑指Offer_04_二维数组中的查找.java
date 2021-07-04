package mid;

/**
 * Created by lwd at 2021/7/4
 *
 * @Description:
 */
public class 剑指Offer_04_二维数组中的查找 {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        int row = matrix.length - 1, col = matrix[0].length - 1;
        int sRow = row, sCol = 0;
        while (sRow >= 0 && sCol <= col) {
            if (matrix[sRow][sCol] > target) {
                sRow--;
            } else if (matrix[sRow][sCol] < target) {
                sCol++;
            } else {
                return true;
            }
        }
        return false;
    }
}
