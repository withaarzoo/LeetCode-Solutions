#include <vector>
#include <queue>
using namespace std;

class Solution
{
public:
    int minCost(vector<vector<int>> &grid)
    {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> directions = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
        vector<vector<int>> cost(m, vector<int>(n, INT_MAX));
        deque<pair<int, int>> dq;
        dq.push_front({0, 0});
        cost[0][0] = 0;

        while (!dq.empty())
        {
            auto [x, y] = dq.front();
            dq.pop_front();

            for (int i = 0; i < 4; ++i)
            {
                int nx = x + directions[i][0], ny = y + directions[i][1];
                int new_cost = cost[x][y] + (grid[x][y] != i + 1);

                if (nx >= 0 && ny >= 0 && nx < m && ny < n && new_cost < cost[nx][ny])
                {
                    cost[nx][ny] = new_cost;
                    if (grid[x][y] == i + 1)
                        dq.push_front({nx, ny});
                    else
                        dq.push_back({nx, ny});
                }
            }
        }
        return cost[m - 1][n - 1];
    }
};
