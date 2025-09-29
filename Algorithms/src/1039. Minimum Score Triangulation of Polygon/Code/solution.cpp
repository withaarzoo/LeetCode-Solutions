class Solution {
public:
    int minScoreTriangulation(vector<int>& values) {
        int n = values.size();
        // dp[i][j] = min score to triangulate polygon from i to j
        vector<vector<int>> dp(n, vector<int>(n, 0));
        
        // len is the number of vertices in the interval
        for (int len = 3; len <= n; ++len) {
            for (int i = 0; i + len - 1 < n; ++i) {
                int j = i + len - 1;
                int best = INT_MAX;
                // choose k as the middle vertex forming triangle (i,k,j)
                for (int k = i + 1; k < j; ++k) {
                    int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                    if (cost < best) best = cost;
                }
                dp[i][j] = best;
            }
        }
        return dp.empty() ? 0 : dp[0][n-1];
    }
};
