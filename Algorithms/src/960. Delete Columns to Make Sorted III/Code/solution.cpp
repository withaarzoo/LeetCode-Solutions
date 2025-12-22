class Solution
{
public:
    int minDeletionSize(vector<string> &strs)
    {
        int n = strs.size();
        int m = strs[0].size();

        // dp[i] = longest valid sequence ending at column i
        vector<int> dp(m, 1);

        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < i; j++)
            {
                bool valid = true;

                // Check all rows
                for (int r = 0; r < n; r++)
                {
                    if (strs[r][j] > strs[r][i])
                    {
                        valid = false;
                        break;
                    }
                }

                if (valid)
                {
                    dp[i] = max(dp[i], dp[j] + 1);
                }
            }
        }

        int keep = *max_element(dp.begin(), dp.end());
        return m - keep;
    }
};
