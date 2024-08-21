class Solution {
    public int strangePrinter(String s) {
        int n = s.length(); // Get the length of the string
        int[][] dp = new int[n][n]; // Create a 2D array to store the minimum turns needed to print each substring

        // Iterate over the string in reverse order to fill the DP table
        for (int i = n - 1; i >= 0; i--) {
            dp[i][i] = 1; // A single character always requires 1 turn to print

            // Consider all substrings starting from index 'i' to 'j'
            for (int j = i + 1; j < n; j++) {
                // Start with the assumption that printing s[j] requires an additional turn
                dp[i][j] = dp[i][j - 1] + 1;

                // Check for all possible partitions of the substring s[i...j]
                for (int k = i; k < j; k++) {
                    // If the character at index 'k' matches the character at index 'j'
                    if (s.charAt(k) == s.charAt(j)) {
                        // We can potentially minimize the number of turns by printing s[j] along with
                        // the segment that ends at 'k'
                        dp[i][j] = Math.min(dp[i][j], dp[i][k] + dp[k + 1][j - 1]);
                    }
                }
            }
        }

        // The minimum number of turns needed to print the entire string s[0...n-1] is
        // stored in dp[0][n-1]
        return dp[0][n - 1];
    }
}
