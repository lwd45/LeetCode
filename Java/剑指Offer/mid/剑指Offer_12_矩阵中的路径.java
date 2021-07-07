package mid;

/**
 * Created by lwd at 2021/7/5
 *
 * @Description:
 */
public class 剑指Offer_12_矩阵中的路径 {
    public boolean exist(char[][] board, String word) {
        if (word == null || word.length() == 0) return true;
        int row = board.length, col = board[0].length;
        for (int i = 0; i < row; ++i) {
            for (int j = 0; j < col; ++j) {
                if (dfs(board, word, i, j, row, col, 0)) return true;
            }
        }
        return false;
    }

    private boolean dfs(char[][] board, String word, int r, int c, int row, int col, int idx) {
        if (r < 0 || c < 0 || r >= row || c >= col || word.charAt(idx) != board[r][c] || board[r][c] == 1) return false;
        if (idx == word.length() - 1) return true;
        board[r][c] = 1;
        boolean res = dfs(board, word, r - 1, c, row, col, idx + 1) || dfs(board, word, r, c - 1, row, col, idx + 1) ||
                dfs(board, word, r + 1, c, row, col, idx + 1) || dfs(board, word, r, c + 1, row, col, idx + 1);
        board[r][c] = word.charAt(idx);
        return res;
    }
}
