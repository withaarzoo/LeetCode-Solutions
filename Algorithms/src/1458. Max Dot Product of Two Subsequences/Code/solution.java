class Solution {
    public int maxDotProduct(int[] nums1, int[] nums2) {
        int n = nums1.length;
        int m = nums2.length;

        int[][] dp = new int[n + 1][m + 1];

        // Initialize with very small values
        for (int i = 0; i <= n; i++) {
            for (int j = 0; j <= m; j++) {
                dp[i][j] = Integer.MIN_VALUE;
            }
        }

        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                int product = nums1[i] * nums2[j];

                int takeBoth = product;
                if (dp[i + 1][j + 1] != Integer.MIN_VALUE) {
                    takeBoth = Math.max(takeBoth, product + dp[i + 1][j + 1]);
                }

                dp[i][j] = Math.max(
                        takeBoth,
                        Math.max(dp[i + 1][j], dp[i][j + 1]));
            }
        }

        return dp[0][0];
    }
}
