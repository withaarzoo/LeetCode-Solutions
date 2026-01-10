class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        n, m = len(s1), len(s2)
        dp = [0] * (m + 1)

        for j in range(m - 1, -1, -1):
            dp[j] = dp[j + 1] + ord(s2[j])

        for i in range(n - 1, -1, -1):
            prev = dp[m]
            dp[m] += ord(s1[i])

            for j in range(m - 1, -1, -1):
                temp = dp[j]
                if s1[i] == s2[j]:
                    dp[j] = prev
                else:
                    dp[j] = min(
                        ord(s1[i]) + dp[j],
                        ord(s2[j]) + dp[j + 1]
                    )
                prev = temp

        return dp[0]
