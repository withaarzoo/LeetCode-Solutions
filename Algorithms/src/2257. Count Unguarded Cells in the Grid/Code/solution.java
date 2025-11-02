import java.util.*;

class Solution {
    public int countUnguarded(int m, int n, int[][] guards, int[][] walls) {
        // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
        int[][] grid = new int[m][n];
        for (int[] w : walls)
            grid[w[0]][w[1]] = 2;
        for (int[] g : guards)
            grid[g[0]][g[1]] = 1;

        int[][] dirs = { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
        for (int[] g : guards) {
            int r = g[0], c = g[1];
            for (int[] dir : dirs) {
                int nr = r + dir[0], nc = c + dir[1];
                while (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                    if (grid[nr][nc] == 2 || grid[nr][nc] == 1)
                        break;
                    if (grid[nr][nc] == 0)
                        grid[nr][nc] = 3;
                    nr += dir[0];
                    nc += dir[1];
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 0)
                    ans++;
        return ans;
    }
}
