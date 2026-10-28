class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[[False] * 1024 for _ in range(n)] for _ in range(m)]
        dp[0][0][grid[0][0]] = True
        for i in range(m):
            for j in range(n):
                if i > 0:
                    for x in range(1024):
                        if dp[i - 1][j][x]:
                            dp[i][j][x ^ grid[i][j]] = True
                if j > 0:
                    for x in range(1024):
                        if dp[i][j - 1][x]:
                            dp[i][j][x ^ grid[i][j]] = True
        for x in range(1024):
            if dp[m - 1][n - 1][x]:
                return x
        return 0