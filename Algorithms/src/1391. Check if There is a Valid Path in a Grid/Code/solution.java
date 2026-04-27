class Solution {
    public boolean hasValidPath(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;

        // Directions: left, right, up, down
        int[][] dirs = {
                { 0, -1 }, // left
                { 0, 1 }, // right
                { -1, 0 }, // up
                { 1, 0 } // down
        };

        // For each street type, which directions it supports.
        // 0 = left, 1 = right, 2 = up, 3 = down
        int[][] streetDirs = {
                {}, // dummy
                { 0, 1 }, // 1: left-right
                { 2, 3 }, // 2: up-down
                { 0, 3 }, // 3: left-down
                { 1, 3 }, // 4: right-down
                { 0, 2 }, // 5: left-up
                { 1, 2 } // 6: right-up
        };

        int[] opposite = { 1, 0, 3, 2 };

        boolean[][] visited = new boolean[m][n];
        java.util.ArrayDeque<int[]> q = new java.util.ArrayDeque<>();
        q.offer(new int[] { 0, 0 });
        visited[0][0] = true;

        while (!q.isEmpty()) {
            int[] cur = q.poll();
            int r = cur[0], c = cur[1];

            if (r == m - 1 && c == n - 1) {
                return true;
            }

            int type = grid[r][c];

            for (int d : streetDirs[type]) {
                int nr = r + dirs[d][0];
                int nc = c + dirs[d][1];

                if (nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc]) {
                    continue;
                }

                int nextType = grid[nr][nc];

                boolean ok = false;
                for (int nd : streetDirs[nextType]) {
                    if (nd == opposite[d]) {
                        ok = true;
                        break;
                    }
                }

                if (ok) {
                    visited[nr][nc] = true;
                    q.offer(new int[] { nr, nc });
                }
            }
        }

        return false;
    }
}