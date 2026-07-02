class Solution
{
public:
    bool findSafeWalk(vector<vector<int>> &grid, int health)
    {
        int m = grid.size(), n = grid[0].size();

        // Store the minimum health lost to reach every cell
        vector<vector<int>> dist(m, vector<int>(n, INT_MAX));

        // Deque used by 0-1 BFS
        deque<pair<int, int>> dq;

        // Starting cost includes the starting cell itself
        dist[0][0] = grid[0][0];
        dq.push_front({0, 0});

        // Four possible movement directions
        int dir[5] = {-1, 0, 1, 0, -1};

        while (!dq.empty())
        {
            auto [x, y] = dq.front();
            dq.pop_front();

            // Try all four neighboring cells
            for (int k = 0; k < 4; k++)
            {
                int nx = x + dir[k];
                int ny = y + dir[k + 1];

                // Ignore cells outside the grid
                if (nx < 0 || ny < 0 || nx >= m || ny >= n)
                    continue;

                // Entering the next cell adds either 0 or 1 cost
                int newCost = dist[x][y] + grid[nx][ny];

                // Update only if this path is better
                if (newCost < dist[nx][ny])
                {
                    dist[nx][ny] = newCost;

                    // Cost 0 goes to the front, cost 1 goes to the back
                    if (grid[nx][ny] == 0)
                        dq.push_front({nx, ny});
                    else
                        dq.push_back({nx, ny});
                }
            }
        }

        // Health must remain at least 1
        return dist[m - 1][n - 1] < health;
    }
};