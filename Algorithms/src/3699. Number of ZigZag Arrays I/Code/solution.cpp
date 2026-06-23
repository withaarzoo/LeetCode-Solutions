class Solution
{
public:
    static constexpr int MOD = 1000000007;

    int zigZagArrays(int n, int l, int r)
    {
        int m = r - l + 1;

        // Length 1: every rank can be chosen once
        vector<int> dp(m, 1);

        for (int len = 2; len <= n; len++)
        {
            // Reversing allows the same prefix-sum logic
            // to act as alternating prefix/suffix transitions
            reverse(dp.begin(), dp.end());

            long long pref = 0;

            for (int i = 0; i < m; i++)
            {
                int old = dp[i]; // Previous DP value

                // New value becomes sum of all earlier values
                dp[i] = pref;

                // Update running prefix sum
                pref = (pref + old) % MOD;
            }
        }

        long long ans = 0;

        // Sum all ending ranks
        for (int x : dp)
        {
            ans = (ans + x) % MOD;
        }

        // Count both starting directions
        return (ans * 2) % MOD;
    }
};