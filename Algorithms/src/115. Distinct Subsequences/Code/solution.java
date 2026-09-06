class Solution {
    public int numDistinct(String s, String t) {
        int m = t.length();

        // dp[j] stores the number of ways to form the first j characters of t.
        // There is exactly one way to form an empty string.
        long[] dp = new long[m + 1];
        dp[0] = 1;

        // Process every character of s.
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            // Traverse backwards so dp[j - 1] is still from the previous state.
            // This prevents using the same character of s more than once.
            for (int j = m; j >= 1; j--) {
                // If the characters match, we can use the current character.
                if (c == t.charAt(j - 1)) {
                    dp[j] += dp[j - 1];
                }
            }
        }

        // The problem guarantees that the answer fits in a 32-bit signed integer.
        return (int) dp[m];
    }
}