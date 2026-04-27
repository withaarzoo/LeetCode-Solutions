class Solution
{
public:
    bool hasValidPath(vector<vector<int>> &grid)
    {
        int m = grid.size();
        int n = grid[0].size();

        // Directions: left, right, up, down
        // Each pair is (dr, dc)
        vector<pair<int, int>> dirs = {
            {0, -1}, // left
            {0, 1},  // right
            {-1, 0}, // up
            {1, 0}   // down
        };

        // For each street type, store which directions it can go to.
        // 0 = left, 1 = right, 2 = up, 3 = down
        vector<vector<int>> streetDirs = {
            {},     // dummy
            {0, 1}, // 1: left-right
            {2, 3}, // 2: up-down
            {0, 3}, // 3: left-down
            {1, 3}, // 4: right-down
            {0, 2}, // 5: left-up
            {1, 2}  // 6: right-up
        };

        // Opposite direction mapping:
        // left <-> right, up <-> down
        vector<int> opposite = {1, 0, 3, 2};

        vector<vector<bool>> visited(m, vector<bool>(n, false));
        queue<pair<int, int>> q;
        q.push({0, 0});
        visited[0][0] = true;

        while (!q.empty())
        {
            auto [r, c] = q.front();
            q.pop();

            if (r == m - 1 && c == n - 1)
            {
                return true;
            }

            int type = grid[r][c];

            for (int d : streetDirs[type])
            {
                int nr = r + dirs[d].first;
                int nc = c + dirs[d].second;

                if (nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc])
                {
                    continue;
                }

                int nextType = grid[nr][nc];

                // Check if the neighbor street connects back to current cell
                bool ok = false;
                for (int nd : streetDirs[nextType])
                {
                    if (nd == opposite[d])
                    {
                        ok = true;
                        break;
                    }
                }

                if (ok)
                {
                    visited[nr][nc] = true;
                    q.push({nr, nc});
                }
            }
        }

        return false;
    }
};