package mid;

/**
 * Created by lwd at 2021/8/16
 *
 * @Description:
 */
public class LeetCode_69_岛屿的最大面积 {
    public int maxAreaOfIsland(int[][] grid) {
        int row = grid.length, col = grid[0].length, max = 0;
        boolean[][] visited = new boolean[row][col];
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (!visited[i][j] && grid[i][j] == 1) {
                    max = Math.max(max, dfs(grid, visited, i, j, row, col));
                }
            }
        }
        return max;
    }

    private int dfs(int[][] grid, boolean[][] visited, int i, int j, int row, int col) {
        if (i < 0 || i >= row || j < 0 || j >= col || visited[i][j] || grid[i][j] == 0) return 0;
        visited[i][j] = true;
        return 1 +
                dfs(grid, visited, i + 1, j, row, col) +
                dfs(grid, visited, i, j + 1, row, col) +
                dfs(grid, visited, i - 1, j, row, col) +
                dfs(grid, visited, i, j - 1, row, col);
    }
}
