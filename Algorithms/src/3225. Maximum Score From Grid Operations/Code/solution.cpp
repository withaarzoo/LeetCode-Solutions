class Solution
{
public:
    long long maximumScore(vector<vector<int>> &grid)
    {
        int n = (int)grid.size();
        if (n == 1)
            return 0;

        // pref[c][k] = sum of first k cells in column c
        vector<vector<long long>> pref(n, vector<long long>(n + 1, 0));
        for (int c = 0; c < n; ++c)
        {
            for (int r = 0; r < n; ++r)
            {
                pref[c][r + 1] = pref[c][r] + grid[r][c];
            }
        }

        const long long NEG = -(1LL << 60);

        // dp[a][b] = best score after processing up to current column,
        // with previous height = a and current height = b.
        vector<vector<long long>> dp(n + 1, vector<long long>(n + 1, NEG));

        // Initialize using the first column (column 0).
        // Its left neighbor is a dummy column of height 0.
        for (int a = 0; a <= n; ++a)
        {
            for (int b = 0; b <= n; ++b)
            {
                dp[a][b] = max(0LL, pref[0][b] - pref[0][a]);
            }
        }

        // Process columns 1..n-1
        for (int col = 1; col < n; ++col)
        {
            vector<vector<long long>> ndp(n + 1, vector<long long>(n + 1, NEG));

            for (int mid = 0; mid <= n; ++mid)
            {
                // q[x] = gain of column 'col' if max(neighbor height) becomes x
                vector<long long> q(n + 1, 0);
                for (int x = 0; x <= n; ++x)
                {
                    q[x] = max(0LL, pref[col][x] - pref[col][mid]);
                }

                // prefixBest[c] = max dp[a][mid] for all a <= c
                vector<long long> prefixBest(n + 1, NEG);
                prefixBest[0] = dp[0][mid];
                for (int a = 1; a <= n; ++a)
                {
                    prefixBest[a] = max(prefixBest[a - 1], dp[a][mid]);
                }

                // suffixBest[c] = max(dp[a][mid] + q[a]) for all a >= c
                vector<long long> suffixBest(n + 2, NEG);
                suffixBest[n] = dp[n][mid] + q[n];
                for (int a = n - 1; a >= 0; --a)
                {
                    suffixBest[a] = max(suffixBest[a + 1], dp[a][mid] + q[a]);
                }

                // For the last real column, the next height is fixed to 0.
                int limit = (col == n - 1 ? 0 : n);

                for (int nxt = 0; nxt <= limit; ++nxt)
                {
                    long long best = NEG;

                    if (prefixBest[nxt] != NEG)
                    {
                        best = max(best, prefixBest[nxt] + q[nxt]);
                    }
                    if (suffixBest[nxt + 1] != NEG)
                    {
                        best = max(best, suffixBest[nxt + 1]);
                    }

                    ndp[mid][nxt] = max(ndp[mid][nxt], best);
                }
            }

            dp.swap(ndp);
        }

        long long ans = 0;
        for (int a = 0; a <= n; ++a)
        {
            for (int b = 0; b <= n; ++b)
            {
                ans = max(ans, dp[a][b]);
            }
        }
        return ans;
    }
};