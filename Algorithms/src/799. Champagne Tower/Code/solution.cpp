class Solution {
public:
    double champagneTower(int poured, int query_row, int query_glass) {
        // Create DP table with 101 rows (safe upper bound)
        vector<vector<double>> dp(101, vector<double>(101, 0.0));
        
        // Pour all champagne into the top glass
        dp[0][0] = poured;
        
        // Simulate the flow row by row
        for (int r = 0; r <= query_row; r++) {
            for (int c = 0; c <= r; c++) {
                // If current glass overflows
                if (dp[r][c] > 1.0) {
                    double overflow = (dp[r][c] - 1.0) / 2.0;
                    
                    // Distribute overflow to next row
                    dp[r + 1][c] += overflow;
                    dp[r + 1][c + 1] += overflow;
                    
                    // Cap current glass to 1
                    dp[r][c] = 1.0;
                }
            }
        }
        
        return min(1.0, dp[query_row][query_glass]);
    }
};
