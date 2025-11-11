class Solution
{
public:
    int findMaxForm(vector<string> &strs, int m, int n)
    {
        // dp[z][o] = max strings using at most z zeros and o ones
        vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));

        for (const string &s : strs)
        {
            int z = 0, o = 0;
            for (char c : s)
                (c == '0') ? ++z : ++o;

            // Go backwards to avoid reusing the same string more than once
            for (int i = m; i >= z; --i)
            {
                for (int j = n; j >= o; --j)
                {
                    dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1);
                }
            }
        }
        return dp[m][n];
    }
};
