class Solution:
    def maximumProfit(self, prices, k):
        n = len(prices)
        NEG = -10**18

        # dp[i][state][t]
        dp = [[[NEG] * (k + 1) for _ in range(3)] for _ in range(n + 1)]

        # Base case
        for t in range(k + 1):
            dp[n][0][t] = 0      # no position is valid
            dp[n][1][t] = NEG   # holding long is invalid
            dp[n][2][t] = NEG   # holding short is invalid

        for i in range(n - 1, -1, -1):
            for t in range(k + 1):
                # state 0: free
                dp[i][0][t] = dp[i + 1][0][t]
                dp[i][0][t] = max(dp[i][0][t],
                                  -prices[i] + dp[i + 1][1][t])  # buy
                dp[i][0][t] = max(dp[i][0][t],
                                   prices[i] + dp[i + 1][2][t])  # short sell

                if t < k:
                    # state 1: holding long
                    dp[i][1][t] = max(
                        dp[i + 1][1][t],                    # hold
                        prices[i] + dp[i + 1][0][t + 1]     # sell
                    )

                    # state 2: holding short
                    dp[i][2][t] = max(
                        dp[i + 1][2][t],                    # hold
                        -prices[i] + dp[i + 1][0][t + 1]    # buy back
                    )

        return dp[0][0][0]
