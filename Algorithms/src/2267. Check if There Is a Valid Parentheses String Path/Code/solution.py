class Solution:
    def hasValidPath(self, grid: list[list[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        if (m + n - 1) % 2 != 0:
            return False
        maxBal = m + n
        dp = [[[False] * maxBal for _ in range(n)] for _ in range(m)]
        start = 1 if grid[0][0] == '(' else -1
        if start < 0:
            return False
        dp[0][0][start] = True
        for i in range(m):
            for j in range(n):
                for bal in range(maxBal):
                    if not dp[i][j][bal]:
                        continue
                    if i + 1 < m:
                        nb = bal + (1 if grid[i + 1][j] == '(' else -1)
                        if 0 <= nb < maxBal:
                            dp[i + 1][j][nb] = True
                    if j + 1 < n:
                        nb = bal + (1 if grid[i][j + 1] == '(' else -1)
                        if 0 <= nb < maxBal:
                            dp[i][j + 1][nb] = True
        return dp[m - 1][n - 1][0]