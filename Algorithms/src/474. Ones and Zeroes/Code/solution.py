from typing import List

class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        # dp[z][o] = max strings using at most z zeros and o ones
        dp = [[0] * (n + 1) for _ in range(m + 1)]

        for s in strs:
            z = s.count('0')
            o = len(s) - z

            # Backwards loops -> 0/1 knapsack
            for i in range(m, z - 1, -1):
                for j in range(n, o - 1, -1):
                    dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1)

        return dp[m][n]
