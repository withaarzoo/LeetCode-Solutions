class Solution {
    public boolean containsCycle(char[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        boolean[][] visited = new boolean[m][n];

        int[] dr = { 1, -1, 0, 0 };
        int[] dc = { 0, 0, 1, -1 };

        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (visited[r][c])
                    continue;

                ArrayDeque<int[]> stack = new ArrayDeque<>();
                stack.addLast(new int[] { r, c, -1, -1 });
                visited[r][c] = true;

                while (!stack.isEmpty()) {
                    int[] cur = stack.removeLast();
                    int cr = cur[0], cc = cur[1], pr = cur[2], pc = cur[3];

                    for (int k = 0; k < 4; k++) {
                        int nr = cr + dr[k];
                        int nc = cc + dc[k];

                        if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                            continue;
                        if (grid[nr][nc] != grid[cr][cc])
                            continue;
                        if (nr == pr && nc == pc)
                            continue;

                        if (visited[nr][nc])
                            return true;

                        visited[nr][nc] = true;
                        stack.addLast(new int[] { nr, nc, cr, cc });
                    }
                }
            }
        }

        return false;
    }
}