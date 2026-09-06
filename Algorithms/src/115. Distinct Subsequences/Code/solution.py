class Solution:
    def numDistinct(self, s: str, t: str) -> int:
        m = len(t)

        # dp[j] stores the number of ways to form the first j characters of t.
        # There is exactly one way to form an empty string.
        dp = [0] * (m + 1)
        dp[0] = 1

        # Process every character of s.
        for c in s:
            # Go backwards so dp[j - 1] still belongs to the previous state.
            # This prevents the current character from being used multiple times.
            for j in range(m, 0, -1):
                # If the characters match, use the current character to extend
                # every subsequence that already forms the previous prefix.
                if c == t[j - 1]:
                    dp[j] += dp[j - 1]

        # dp[m] contains the number of ways to form all of t.
        return dp[m]