#include <vector>
#include <algorithm>
using namespace std;

class Solution
{
public:
    int maxMoves(vector<vector<int>> &grid)
    {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> dp(m, vector<int>(n, 0));
        int maxMoves = 0;

        for (int col = n - 2; col >= 0; --col)
        {
            for (int row = 0; row < m; ++row)
            {
                if (row > 0 && grid[row][col] < grid[row - 1][col + 1])
                {
                    dp[row][col] = max(dp[row][col], dp[row - 1][col + 1] + 1);
                }
                if (grid[row][col] < grid[row][col + 1])
                {
                    dp[row][col] = max(dp[row][col], dp[row][col + 1] + 1);
                }
                if (row < m - 1 && grid[row][col] < grid[row + 1][col + 1])
                {
                    dp[row][col] = max(dp[row][col], dp[row + 1][col + 1] + 1);
                }
            }
        }

        for (int row = 0; row < m; ++row)
        {
            maxMoves = max(maxMoves, dp[row][0]);
        }
        return maxMoves;
    }
};