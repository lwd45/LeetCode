package mid;

import sun.security.util.Length;

/**
 * Created by lwd at 2021/8/13
 *
 * @Description:
 */
public class LeetCode_73_矩阵置零 {
    public void setZeroes(int[][] matrix) {
        boolean zero_row_have_0 = false, zero_col_have_0 = false;
        int row = matrix.length, col = matrix[0].length;

        for (int i = 0; i < col; i++) {
            if (matrix[0][i] == 0) {
                zero_row_have_0 = true;
                break;
            }
        }
        for (int j = 0; j < row; j++) {
            if (matrix[0][j] == 0) {
                zero_col_have_0 = true;
                break;
            }
        }

        for (int i = 1; i < row; i++) {
            for (int j = 1; j < col; j++) {
                if (matrix[i][j] == 0) {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }

        for (int i = 1; i < row; i++) {
            for (int j = 1; j < col; j++) {
                if (matrix[i][0] == 0 || matrix[0][j] == 0) {
                    matrix[i][j] = 0;
                }
            }
        }
        if (zero_row_have_0) {
            for (int i = 0; i < col; i++) {
                matrix[0][i] = 0;
            }
        }
        if (zero_col_have_0) {
            for (int i = 0; i < row; i++) {
                matrix[i][0] = 0;
            }
        }
    }
}
