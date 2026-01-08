class Solution {
public:
    int maxDotProduct(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        int m = nums2.size();
        
        // dp[i][j] = max dot product using nums1[i:] and nums2[j:]
        vector<vector<int>> dp(n + 1, vector<int>(m + 1, INT_MIN));
        
        // Fill DP table from bottom-right to top-left
        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                int product = nums1[i] * nums2[j];
                
                // Option 1: take both elements
                int takeBoth = product;
                if (dp[i + 1][j + 1] != INT_MIN)
                    takeBoth = max(takeBoth, product + dp[i + 1][j + 1]);
                
                // Option 2 & 3: skip one element
                dp[i][j] = max({
                    takeBoth,
                    dp[i + 1][j],
                    dp[i][j + 1]
                });
            }
        }
        
        return dp[0][0];
    }
};
