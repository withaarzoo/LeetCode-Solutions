from typing import List

class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        MOD = 10**9 + 7
        m, n = len(grid), len(grid[0])

        # prev[j][r] = paths to (i-1, j) with sum % k == r
        # cur[j][r]  = paths to (i,   j) with sum % k == r
        prev = [[0] * k for _ in range(n)]
        cur  = [[0] * k for _ in range(n)]

        for i in range(m):
            # reset current row
            for j in range(n):
                for r in range(k):
                    cur[j][r] = 0

            for j in range(n):
                val = grid[i][j] % k

                # starting cell
                if i == 0 and j == 0:
                    cur[0][val] = 1
                    continue

                # from top
                if i > 0:
                    for r in range(k):
                        if prev[j][r] == 0:
                            continue
                        nr = (r + val) % k
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD

                # from left
                if j > 0:
                    for r in range(k):
                        if cur[j - 1][r] == 0:
                            continue
                        nr = (r + val) % k
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD

            prev, cur = cur, prev

        return prev[n - 1][0]
