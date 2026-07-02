class Solution {
    public boolean findSafeWalk(List<List<Integer>> grid, int health) {
        int m = grid.size();
        int n = grid.get(0).size();

        // Store minimum health lost for every cell
        int[][] dist = new int[m][n];
        for (int i = 0; i < m; i++) {
            java.util.Arrays.fill(dist[i], Integer.MAX_VALUE);
        }

        // Deque for 0-1 BFS
        java.util.ArrayDeque<int[]> dq = new java.util.ArrayDeque<>();

        // Starting cost includes the starting cell
        dist[0][0] = grid.get(0).get(0);
        dq.offerFirst(new int[] { 0, 0 });

        int[] dir = { -1, 0, 1, 0, -1 };

        while (!dq.isEmpty()) {
            int[] cur = dq.pollFirst();
            int x = cur[0];
            int y = cur[1];

            // Visit all four neighbors
            for (int k = 0; k < 4; k++) {
                int nx = x + dir[k];
                int ny = y + dir[k + 1];

                // Skip invalid positions
                if (nx < 0 || ny < 0 || nx >= m || ny >= n)
                    continue;

                // Add the cost of entering the next cell
                int newCost = dist[x][y] + grid.get(nx).get(ny);

                // Keep only the best cost
                if (newCost < dist[nx][ny]) {
                    dist[nx][ny] = newCost;

                    // Weight 0 goes to the front, weight 1 goes to the back
                    if (grid.get(nx).get(ny) == 0)
                        dq.offerFirst(new int[] { nx, ny });
                    else
                        dq.offerLast(new int[] { nx, ny });
                }
            }
        }

        // Final health must be positive
        return dist[m - 1][n - 1] < health;
    }
}