class Solution
{
public:
    bool containsCycle(vector<vector<char>> &grid)
    {
        int m = grid.size();
        int n = grid[0].size();
        vector<vector<bool>> visited(m, vector<bool>(n, false));

        int dr[4] = {1, -1, 0, 0};
        int dc[4] = {0, 0, 1, -1};

        for (int r = 0; r < m; ++r)
        {
            for (int c = 0; c < n; ++c)
            {
                if (visited[r][c])
                    continue;

                // stack node: {row, col, parentRow, parentCol}
                vector<array<int, 4>> st;
                st.push_back({r, c, -1, -1});
                visited[r][c] = true;

                while (!st.empty())
                {
                    auto cur = st.back();
                    st.pop_back();

                    int cr = cur[0], cc = cur[1];
                    int pr = cur[2], pc = cur[3];

                    for (int k = 0; k < 4; ++k)
                    {
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
                        st.push_back({nr, nc, cr, cc});
                    }
                }
            }
        }

        return false;
    }
};