from typing import List

class Solution:
    def minScoreTriangulation(self, values: List[int]) -> int:
        n = len(values)
        if n < 3:
            return 0
        # dp[i][j] = min score for sub-polygon from i to j
        dp = [[0] * n for _ in range(n)]
        
        # compute for all lengths starting from 3
        for length in range(3, n + 1):
            for i in range(0, n - length + 1):
                j = i + length - 1
                best = float('inf')
                for k in range(i + 1, j):
                    cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j]
                    if cost < best:
                        best = cost
                dp[i][j] = best
        return dp[0][n-1]
