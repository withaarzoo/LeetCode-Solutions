class Solution {
    public String findTheString(int[][] lcp) {
        int n = lcp.length;

        int[] group = new int[n];
        Arrays.fill(group, -1);

        int curGroup = 0;

        // Build groups
        for (int i = 0; i < n; i++) {
            if (group[i] == -1) {
                if (curGroup == 26)
                    return "";

                group[i] = curGroup++;

                for (int j = i + 1; j < n; j++) {
                    if (lcp[i][j] > 0) {
                        group[j] = group[i];
                    }
                }
            }
        }

        // Build answer string
        char[] ans = new char[n];
        for (int i = 0; i < n; i++) {
            ans[i] = (char) ('a' + group[i]);
        }

        // Verify using DP
        int[][] dp = new int[n + 1][n + 1];

        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                if (ans[i] == ans[j]) {
                    dp[i][j] = 1 + dp[i + 1][j + 1];
                }
            }
        }

        // Compare matrices
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (dp[i][j] != lcp[i][j]) {
                    return "";
                }
            }
        }

        return new String(ans);
    }
}