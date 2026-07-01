class Solution {
    public int maximumSafenessFactor(List<List<Integer>> grid) {

        int n = grid.size();

        int[][] dist = new int[n][n];
        for (int[] row : dist)
            Arrays.fill(row, -1);

        Queue<int[]> queue = new LinkedList<>();

        // Start BFS from every thief
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (grid.get(i).get(j) == 1) {
                    dist[i][j] = 0;
                    queue.offer(new int[] { i, j });
                }
            }
        }

        int[] dx = { -1, 0, 1, 0 };
        int[] dy = { 0, 1, 0, -1 };

        // Multi-source BFS
        while (!queue.isEmpty()) {
            int[] cur = queue.poll();

            for (int k = 0; k < 4; k++) {
                int nx = cur[0] + dx[k];
                int ny = cur[1] + dy[k];

                if (nx < 0 || ny < 0 || nx >= n || ny >= n || dist[nx][ny] != -1)
                    continue;

                dist[nx][ny] = dist[cur[0]][cur[1]] + 1;
                queue.offer(new int[] { nx, ny });
            }
        }

        int left = 0;
        int right = 2 * n;
        int ans = 0;

        while (left <= right) {

            int mid = left + (right - left) / 2;

            if (canReach(dist, mid, n)) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return ans;
    }

    private boolean canReach(int[][] dist, int limit, int n) {

        if (dist[0][0] < limit || dist[n - 1][n - 1] < limit)
            return false;

        Queue<int[]> queue = new LinkedList<>();
        boolean[][] vis = new boolean[n][n];

        queue.offer(new int[] { 0, 0 });
        vis[0][0] = true;

        int[] dx = { -1, 0, 1, 0 };
        int[] dy = { 0, 1, 0, -1 };

        while (!queue.isEmpty()) {

            int[] cur = queue.poll();

            if (cur[0] == n - 1 && cur[1] == n - 1)
                return true;

            for (int k = 0; k < 4; k++) {

                int nx = cur[0] + dx[k];
                int ny = cur[1] + dy[k];

                if (nx < 0 || ny < 0 || nx >= n || ny >= n)
                    continue;

                if (vis[nx][ny] || dist[nx][ny] < limit)
                    continue;

                vis[nx][ny] = true;
                queue.offer(new int[] { nx, ny });
            }
        }

        return false;
    }
}