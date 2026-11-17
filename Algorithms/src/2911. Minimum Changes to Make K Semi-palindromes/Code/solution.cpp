class Solution {
public:
    int minimumChanges(string s, int k) {
        int n = s.size();
        const int INF = 1e9;
        vector<vector<int>> cost(n, vector<int>(n, INF));
        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                int L = j - i + 1;
                int best = INF;
                for (int d = 1; d < L; ++d) {
                    if (L % d != 0) continue;
                    int cur = 0;
                    int m = L / d;
                    for (int r = 0; r < d; ++r) {
                        for (int t = 0; t < m / 2; ++t) {
                            if (s[i + r + t * d] != s[i + r + (m - 1 - t) * d])
                                ++cur;
                        }
                    }
                    best = min(best, cur);
                }
                cost[i][j] = best;
            }
        }
        vector<vector<int>> dp(n + 1, vector<int>(k + 1, INF));
        dp[0][0] = 0;
        for (int i = 1; i <= n; ++i) {
            for (int p = 1; p <= k; ++p) {
                for (int prev = 0; prev <= i - 2; ++prev) {
                    if (dp[prev][p - 1] != INF && cost[prev][i - 1] != INF)
                        dp[i][p] = min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
                }
            }
        }
        return dp[n][k];
    }
};