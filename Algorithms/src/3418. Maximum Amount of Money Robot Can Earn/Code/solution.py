class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        m, n = len(coins), len(coins[0])
        NEG = -10**9

        # dp[i][j][k] = max money at (i,j) using k neutralizations
        dp = [[[NEG] * 3 for _ in range(n)] for _ in range(m)]

        # Starting cell
        if coins[0][0] >= 0:
            dp[0][0][0] = coins[0][0]
        else:
            dp[0][0][0] = coins[0][0]
            dp[0][0][1] = 0

        for i in range(m):
            for j in range(n):
                for k in range(3):
                    if dp[i][j][k] == NEG:
                        continue

                    # Move Down
                    if i + 1 < m:
                        val = coins[i + 1][j]

                        # Take normally
                        dp[i + 1][j][k] = max(dp[i + 1][j][k], dp[i][j][k] + val)

                        # Neutralize if negative
                        if val < 0 and k < 2:
                            dp[i + 1][j][k + 1] = max(dp[i + 1][j][k + 1], dp[i][j][k])

                    # Move Right
                    if j + 1 < n:
                        val = coins[i][j + 1]

                        # Take normally
                        dp[i][j + 1][k] = max(dp[i][j + 1][k], dp[i][j][k] + val)

                        # Neutralize if negative
                        if val < 0 and k < 2:
                            dp[i][j + 1][k + 1] = max(dp[i][j + 1][k + 1], dp[i][j][k])

        return max(dp[m - 1][n - 1])