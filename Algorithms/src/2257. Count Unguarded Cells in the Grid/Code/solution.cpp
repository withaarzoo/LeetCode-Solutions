#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int countUnguarded(int m, int n, vector<vector<int>> &guards, vector<vector<int>> &walls)
    {
        // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
        vector<vector<int>> grid(m, vector<int>(n, 0));
        for (auto &w : walls)
            grid[w[0]][w[1]] = 2;
        for (auto &g : guards)
            grid[g[0]][g[1]] = 1;

        // Directions: up, down, left, right
        const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

        for (auto &g : guards)
        {
            int r = g[0], c = g[1];
            for (int d = 0; d < 4; ++d)
            {
                int nr = r + dirs[d][0], nc = c + dirs[d][1];
                while (nr >= 0 && nr < m && nc >= 0 && nc < n)
                {
                    if (grid[nr][nc] == 2 || grid[nr][nc] == 1)
                        break; // wall or guard stops view
                    if (grid[nr][nc] == 0)
                        grid[nr][nc] = 3; // mark as guarded
                    nr += dirs[d][0];
                    nc += dirs[d][1];
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 0)
                    ++ans; // unguarded and unoccupied
        return ans;
    }
};
