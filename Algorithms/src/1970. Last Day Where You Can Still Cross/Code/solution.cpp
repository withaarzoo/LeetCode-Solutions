class Solution
{
public:
    int latestDayToCross(int row, int col, vector<vector<int>> &cells)
    {
        int n = row * col;
        int top = n, bottom = n + 1;

        vector<int> parent(n + 2), rank(n + 2, 0);
        vector<vector<bool>> grid(row, vector<bool>(col, false));

        for (int i = 0; i < n + 2; i++)
            parent[i] = i;

        function<int(int)> find = [&](int x)
        {
            if (parent[x] != x)
                parent[x] = find(parent[x]);
            return parent[x];
        };

        auto unite = [&](int a, int b)
        {
            a = find(a);
            b = find(b);
            if (a == b)
                return;
            if (rank[a] < rank[b])
                swap(a, b);
            parent[b] = a;
            if (rank[a] == rank[b])
                rank[a]++;
        };

        int dr[4] = {1, -1, 0, 0};
        int dc[4] = {0, 0, 1, -1};

        for (int d = n - 1; d >= 0; d--)
        {
            int r = cells[d][0] - 1;
            int c = cells[d][1] - 1;
            grid[r][c] = true;
            int id = r * col + c;

            if (r == 0)
                unite(id, top);
            if (r == row - 1)
                unite(id, bottom);

            for (int k = 0; k < 4; k++)
            {
                int nr = r + dr[k];
                int nc = c + dc[k];
                if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc])
                {
                    unite(id, nr * col + nc);
                }
            }

            if (find(top) == find(bottom))
                return d;
        }
        return 0;
    }
};
