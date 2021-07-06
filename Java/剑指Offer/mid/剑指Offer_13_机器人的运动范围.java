package mid;

/**
 * Created by lwd at 2021/7/6
 *
 * @Description:
 */
public class 剑指Offer_13_机器人的运动范围 {
    boolean[][] visited;

    public int movingCount(int m, int n, int k) {
        visited = new boolean[m][n];
        return dfs(0, 0, m, n, k);
    }

    private int dfs(int r, int c, int m, int n, int k) {
        if (r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || !canMove(r, c, k)) return 0;

        visited[r][c] = true;
        return 1 + dfs(r + 1, c, m, n, k) + dfs(r, c + 1, m, n, k);
    }

    private boolean canMove(int c, int r, int k) {
        int sum = 0;
        while (c > 0) {
            sum += c % 10;
            c = c / 10;
        }
        while (r > 0) {
            sum += r % 10;
            r = r / 10;
        }
        return sum <= k;
    }

}
