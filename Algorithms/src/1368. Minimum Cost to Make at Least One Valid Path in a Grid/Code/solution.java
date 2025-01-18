import java.util.*;

class Solution {
    public int minCost(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[][] directions = { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } };
        int[][] cost = new int[m][n];
        for (int[] row : cost)
            Arrays.fill(row, Integer.MAX_VALUE);

        Deque<int[]> dq = new ArrayDeque<>();
        dq.offerFirst(new int[] { 0, 0 });
        cost[0][0] = 0;

        while (!dq.isEmpty()) {
            int[] cell = dq.pollFirst();
            int x = cell[0], y = cell[1];

            for (int i = 0; i < 4; ++i) {
                int nx = x + directions[i][0];
                int ny = y + directions[i][1];
                int newCost = cost[x][y] + (grid[x][y] != i + 1);

                if (nx >= 0 && ny >= 0 && nx < m && ny < n && newCost < cost[nx][ny]) {
                    cost[nx][ny] = newCost;
                    if (grid[x][y] == i + 1)
                        dq.offerFirst(new int[] { nx, ny });
                    else
                        dq.offerLast(new int[] { nx, ny });
                }
            }
        }
        return cost[m - 1][n - 1];
    }
}
