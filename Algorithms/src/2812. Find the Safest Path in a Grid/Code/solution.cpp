class Solution
{
public:
    int maximumSafenessFactor(vector<vector<int>> &grid)
    {
        int n = grid.size();

        // Distance of every cell from the nearest thief
        vector<vector<int>> dist(n, vector<int>(n, -1));

        queue<pair<int, int>> q;

        // Push every thief into the queue
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (grid[i][j] == 1)
                {
                    dist[i][j] = 0;
                    q.push({i, j});
                }
            }
        }

        int dir[5] = {-1, 0, 1, 0, -1};

        // Multi-source BFS
        while (!q.empty())
        {
            auto [x, y] = q.front();
            q.pop();

            for (int k = 0; k < 4; k++)
            {
                int nx = x + dir[k];
                int ny = y + dir[k + 1];

                if (nx < 0 || ny < 0 || nx >= n || ny >= n || dist[nx][ny] != -1)
                    continue;

                // First visit always gives the shortest distance
                dist[nx][ny] = dist[x][y] + 1;
                q.push({nx, ny});
            }
        }

        // Check if a path exists with safeness >= limit
        auto canReach = [&](int limit)
        {
            if (dist[0][0] < limit || dist[n - 1][n - 1] < limit)
                return false;

            vector<vector<int>> vis(n, vector<int>(n, 0));
            queue<pair<int, int>> bfs;

            bfs.push({0, 0});
            vis[0][0] = 1;

            while (!bfs.empty())
            {
                auto [x, y] = bfs.front();
                bfs.pop();

                if (x == n - 1 && y == n - 1)
                    return true;

                for (int k = 0; k < 4; k++)
                {
                    int nx = x + dir[k];
                    int ny = y + dir[k + 1];

                    if (nx < 0 || ny < 0 || nx >= n || ny >= n)
                        continue;

                    if (vis[nx][ny])
                        continue;

                    // Only move through safe enough cells
                    if (dist[nx][ny] < limit)
                        continue;

                    vis[nx][ny] = 1;
                    bfs.push({nx, ny});
                }
            }

            return false;
        };

        int left = 0;
        int right = 2 * n;
        int ans = 0;

        // Binary search on the answer
        while (left <= right)
        {
            int mid = left + (right - left) / 2;

            if (canReach(mid))
            {
                ans = mid;
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }

        return ans;
    }
};