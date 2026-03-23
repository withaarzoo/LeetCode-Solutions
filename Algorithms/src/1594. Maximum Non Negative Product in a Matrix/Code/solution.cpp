class Solution {
public:
    int maxProductPath(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        const long long MOD = 1e9 + 7;

        vector<vector<long long>> maxDp(m, vector<long long>(n));
        vector<vector<long long>> minDp(m, vector<long long>(n));

        maxDp[0][0] = minDp[0][0] = grid[0][0];

        // First column
        for (int i = 1; i < m; i++) {
            maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0];
        }

        // First row
        for (int j = 1; j < n; j++) {
            maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j];
        }

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                long long val = grid[i][j];

                long long a = maxDp[i-1][j] * val;
                long long b = minDp[i-1][j] * val;
                long long c = maxDp[i][j-1] * val;
                long long d = minDp[i][j-1] * val;

                maxDp[i][j] = max({a, b, c, d});
                minDp[i][j] = min({a, b, c, d});
            }
        }

        long long res = maxDp[m-1][n-1];
        if (res < 0) return -1;
        return res % MOD;
    }
};