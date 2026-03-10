class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:

        MOD = 10**9 + 7

        dp = [[[0,0] for _ in range(one+1)] for _ in range(zero+1)]

        for i in range(1, min(zero,limit)+1):
            dp[i][0][0] = 1

        for j in range(1, min(one,limit)+1):
            dp[0][j][1] = 1

        for i in range(1, zero+1):
            for j in range(1, one+1):

                over0 = dp[i-limit-1][j][1] if i-limit-1 >= 0 else 0
                over1 = dp[i][j-limit-1][0] if j-limit-1 >= 0 else 0

                dp[i][j][0] = (
                    dp[i-1][j][0] + dp[i-1][j][1] - over0
                ) % MOD

                dp[i][j][1] = (
                    dp[i][j-1][0] + dp[i][j-1][1] - over1
                ) % MOD

        return (dp[zero][one][0] + dp[zero][one][1]) % MOD