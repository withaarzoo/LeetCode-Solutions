import java.util.PriorityQueue;

class Solution {
    public int trapRainWater(int[][] heightMap) {
        int m = heightMap.length, n = heightMap[0].length;
        if (m < 3 || n < 3)
            return 0;

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        boolean[][] visited = new boolean[m][n];

        // Add all boundary cells
        for (int i = 0; i < m; i++) {
            pq.offer(new int[] { heightMap[i][0], i, 0 });
            pq.offer(new int[] { heightMap[i][n - 1], i, n - 1 });
            visited[i][0] = visited[i][n - 1] = true;
        }
        for (int j = 0; j < n; j++) {
            pq.offer(new int[] { heightMap[0][j], 0, j });
            pq.offer(new int[] { heightMap[m - 1][j], m - 1, j });
            visited[0][j] = visited[m - 1][j] = true;
        }

        int result = 0;
        int[][] directions = { { 0, 1 }, { 1, 0 }, { 0, -1 }, { -1, 0 } };

        while (!pq.isEmpty()) {
            int[] cell = pq.poll();
            int height = cell[0], x = cell[1], y = cell[2];

            for (int[] dir : directions) {
                int nx = x + dir[0], ny = y + dir[1];
                if (nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny]) {
                    result += Math.max(0, height - heightMap[nx][ny]);
                    pq.offer(new int[] { Math.max(height, heightMap[nx][ny]), nx, ny });
                    visited[nx][ny] = true;
                }
            }
        }

        return result;
    }
}
