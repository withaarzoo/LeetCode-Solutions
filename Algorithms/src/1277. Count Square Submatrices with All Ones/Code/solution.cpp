#include <vector>
#include <algorithm>

class Solution
{
public:
    int countSquares(std::vector<std::vector<int>> &matrix)
    {
        int m = matrix.size();
        int n = matrix[0].size();
        std::vector<std::vector<int>> dp(m, std::vector<int>(n, 0));
        int totalSquares = 0;

        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (matrix[i][j] == 1)
                {
                    if (i == 0 || j == 0)
                    {
                        dp[i][j] = 1;
                    }
                    else
                    {
                        dp[i][j] = std::min({dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1]}) + 1;
                    }
                    totalSquares += dp[i][j];
                }
            }
        }

        return totalSquares;
    }
};