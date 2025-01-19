#include <vector>
#include <queue>
#include <tuple>
using namespace std;

class Solution
{
public:
    int trapRainWater(vector<vector<int>> &heightMap)
    {
        int m = heightMap.size(), n = heightMap[0].size();
        if (m < 3 || n < 3)
            return 0;

        priority_queue<tuple<int, int, int>, vector<tuple<int, int, int>>, greater<>> pq;
        vector<vector<bool>> visited(m, vector<bool>(n, false));

        // Add all boundary cells to the priority queue
        for (int i = 0; i < m; ++i)
        {
            pq.emplace(heightMap[i][0], i, 0);
            pq.emplace(heightMap[i][n - 1], i, n - 1);
            visited[i][0] = visited[i][n - 1] = true;
        }
        for (int j = 0; j < n; ++j)
        {
            pq.emplace(heightMap[0][j], 0, j);
            pq.emplace(heightMap[m - 1][j], m - 1, j);
            visited[0][j] = visited[m - 1][j] = true;
        }

        int result = 0, directions[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

        while (!pq.empty())
        {
            auto [height, x, y] = pq.top();
            pq.pop();

            for (auto &dir : directions)
            {
                int nx = x + dir[0], ny = y + dir[1];
                if (nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny])
                {
                    result += max(0, height - heightMap[nx][ny]);
                    pq.emplace(max(height, heightMap[nx][ny]), nx, ny);
                    visited[nx][ny] = true;
                }
            }
        }

        return result;
    }
};
