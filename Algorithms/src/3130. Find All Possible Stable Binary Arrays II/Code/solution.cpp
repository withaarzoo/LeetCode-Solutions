class Solution
{
public:
    int numberOfStableArrays(int zero, int one, int limit)
    {
        const int MOD = 1e9 + 7;

        vector<vector<array<long long, 2>>> dp(
            zero + 1, vector<array<long long, 2>>(one + 1, {0, 0}));

        // base cases
        for (int i = 1; i <= min(zero, limit); i++)
            dp[i][0][0] = 1;

        for (int j = 1; j <= min(one, limit); j++)
            dp[0][j][1] = 1;

        for (int i = 1; i <= zero; i++)
        {
            for (int j = 1; j <= one; j++)
            {

                long long over0 = (i - limit - 1 >= 0) ? dp[i - limit - 1][j][1] : 0;
                long long over1 = (j - limit - 1 >= 0) ? dp[i][j - limit - 1][0] : 0;

                dp[i][j][0] =
                    (dp[i - 1][j][0] + dp[i - 1][j][1] - over0 + MOD) % MOD;

                dp[i][j][1] =
                    (dp[i][j - 1][0] + dp[i][j - 1][1] - over1 + MOD) % MOD;
            }
        }

        return (dp[zero][one][0] + dp[zero][one][1]) % MOD;
    }
};