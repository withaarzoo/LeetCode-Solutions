class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int n = triangle.size();
        // dp initialized with last row (copy)
        vector<int> dp(triangle.back());
        // process rows from bottom-1 up to 0
        for (int i = n - 2; i >= 0; --i) {
            for (int j = 0; j <= i; ++j) {
                // choose the smaller of the two adjacent values below
                dp[j] = triangle[i][j] + min(dp[j], dp[j + 1]);
            }
        }
        return dp[0];
    }
};
