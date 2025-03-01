class Solution
{
public:
    long long gridGame(vector<vector<int>> &grid)
    {
        int n = grid[0].size();
        vector<long long> topSuffix(n), bottomPrefix(n);

        // Calculate suffix sum for the top row
        topSuffix[n - 1] = grid[0][n - 1];
        for (int i = n - 2; i >= 0; --i)
        {
            topSuffix[i] = topSuffix[i + 1] + grid[0][i];
        }

        // Calculate prefix sum for the bottom row
        bottomPrefix[0] = grid[1][0];
        for (int i = 1; i < n; ++i)
        {
            bottomPrefix[i] = bottomPrefix[i - 1] + grid[1][i];
        }

        // Find the minimum maximum points Robot 2 can collect
        long long result = LLONG_MAX;
        for (int i = 0; i < n; ++i)
        {
            long long top = (i + 1 < n) ? topSuffix[i + 1] : 0;
            long long bottom = (i > 0) ? bottomPrefix[i - 1] : 0;
            result = min(result, max(top, bottom));
        }

        return result;
    }
};
