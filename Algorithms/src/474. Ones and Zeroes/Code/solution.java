class Solution {
    public int findMaxForm(String[] strs, int m, int n) {
        // dp[z][o] = max strings using at most z zeros and o ones
        int[][] dp = new int[m + 1][n + 1];

        for (String s : strs) {
            int z = 0, o = 0;
            for (char c : s.toCharArray()) {
                if (c == '0')
                    z++;
                else
                    o++;
            }

            // Backwards iteration for 0/1 knapsack
            for (int i = m; i >= z; i--) {
                for (int j = n; j >= o; j--) {
                    dp[i][j] = Math.max(dp[i][j], dp[i - z][j - o] + 1);
                }
            }
        }
        return dp[m][n];
    }
}
