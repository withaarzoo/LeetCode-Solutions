from typing import List
class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        MOD = 10**9 + 7
        landR = [[[0, 0] for _ in range(n)] for _ in range(m)]
        landC = [[[0, 0] for _ in range(n)] for _ in range(m)]
        for x in range(m - 1, -1, -1):
            for y in range(n - 1, -1, -1):
                for d in range(2):
                    if grid[x][y] == 0:
                        landR[x][y][d] = x
                        landC[x][y][d] = y
                    else:
                        nx = x + 1 if d == 0 else x
                        ny = y if d == 0 else y + 1
                        nd = 1 - d
                        if nx < 0 or nx >= m or ny < 0 or ny >= n:
                            landR[x][y][d] = -1
                            landC[x][y][d] = -1
                        else:
                            landR[x][y][d] = landR[nx][ny][nd]
                            landC[x][y][d] = landC[nx][ny][nd]
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = 1
        for i in range(m):
            for j in range(n):
                if dp[i][j] == 0:
                    continue
                if j + 1 < n:
                    nr = landR[i][j + 1][0]
                    nc = landC[i][j + 1][0]
                    if nr != -1:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                if i + 1 < m:
                    nr = landR[i + 1][j][1]
                    nc = landC[i + 1][j][1]
                    if nr != -1:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
        return dp[m - 1][n - 1]