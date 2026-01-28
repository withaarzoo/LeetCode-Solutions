from typing import List

class Solution:
    def minCost(self, grid: List[List[int]], k: int) -> int:
        INF = 10**18
        m, n = len(grid), len(grid[0])
        # base dp (0 teleports)
        dp = [[INF]*n for _ in range(m)]
        dp[0][0] = 0
        for i in range(m):
            for j in range(n):
                if i > 0:
                    dp[i][j] = min(dp[i][j], dp[i-1][j] + grid[i][j])
                if j > 0:
                    dp[i][j] = min(dp[i][j], dp[i][j-1] + grid[i][j])

        # prepare cells sorted by value descending
        cells = []
        for i in range(m):
            for j in range(n):
                cells.append((grid[i][j], i, j))
        cells.sort(reverse=True)  # desc by value

        for _ in range(k):
            # compute start costs after an extra teleport (or none)
            start = [[INF]*n for _ in range(m)]
            running_min = INF
            idx = 0
            L = len(cells)
            while idx < L:
                val = cells[idx][0]
                j = idx
                min_group = INF
                while j < L and cells[j][0] == val:
                    _, ii, jj = cells[j]
                    min_group = min(min_group, dp[ii][jj])
                    j += 1
                running_min = min(running_min, min_group)
                for p in range(idx, j):
                    _, ii, jj = cells[p]
                    start[ii][jj] = min(dp[ii][jj], running_min)
                idx = j

            # propagate normal right/down moves from start
            dp2 = [[INF]*n for _ in range(m)]
            for i in range(m):
                for j in range(n):
                    dp2[i][j] = min(dp2[i][j], start[i][j])
                    if dp2[i][j] < INF:
                        if i+1 < m:
                            dp2[i+1][j] = min(dp2[i+1][j], dp2[i][j] + grid[i+1][j])
                        if j+1 < n:
                            dp2[i][j+1] = min(dp2[i][j+1], dp2[i][j] + grid[i][j+1])
            dp = dp2

        return dp[m-1][n-1]
