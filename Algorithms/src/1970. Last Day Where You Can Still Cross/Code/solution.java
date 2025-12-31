class Solution {
    public int latestDayToCross(int row, int col, int[][] cells) {
        int n = row * col;
        int top = n, bottom = n + 1;

        int[] parent = new int[n + 2];
        int[] rank = new int[n + 2];
        boolean[][] grid = new boolean[row][col];

        for (int i = 0; i < n + 2; i++)
            parent[i] = i;

        int[] dr = { 1, -1, 0, 0 };
        int[] dc = { 0, 0, 1, -1 };

        for (int d = n - 1; d >= 0; d--) {
            int r = cells[d][0] - 1;
            int c = cells[d][1] - 1;
            grid[r][c] = true;
            int id = r * col + c;

            if (r == 0)
                union(id, top, parent, rank);
            if (r == row - 1)
                union(id, bottom, parent, rank);

            for (int k = 0; k < 4; k++) {
                int nr = r + dr[k];
                int nc = c + dc[k];
                if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc]) {
                    union(id, nr * col + nc, parent, rank);
                }
            }

            if (find(top, parent) == find(bottom, parent))
                return d;
        }
        return 0;
    }

    private int find(int x, int[] parent) {
        if (parent[x] != x)
            parent[x] = find(parent[x], parent);
        return parent[x];
    }

    private void union(int a, int b, int[] parent, int[] rank) {
        a = find(a, parent);
        b = find(b, parent);
        if (a == b)
            return;
        if (rank[a] < rank[b])
            parent[a] = b;
        else {
            parent[b] = a;
            if (rank[a] == rank[b])
                rank[a]++;
        }
    }
}
