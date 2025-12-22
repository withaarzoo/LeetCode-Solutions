class Solution:
    def minDeletionSize(self, strs):
        n = len(strs)
        m = len(strs[0])

        dp = [1] * m

        for i in range(m):
            for j in range(i):
                valid = True
                for r in range(n):
                    if strs[r][j] > strs[r][i]:
                        valid = False
                        break
                if valid:
                    dp[i] = max(dp[i], dp[j] + 1)

        return m - max(dp)
