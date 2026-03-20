from typing import List

class Solution:
    def minAbsDiff(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m = len(grid)
        n = len(grid[0])
        ans = [[0] * (n - k + 1) for _ in range(m - k + 1)]

        for i in range(m - k + 1):
            for j in range(n - k + 1):
                vals = []

                # Collect all values from the current k x k submatrix
                for r in range(i, i + k):
                    for c in range(j, j + k):
                        vals.append(grid[r][c])

                vals.sort()

                best = float('inf')

                # Check only consecutive different values
                for x in range(1, len(vals)):
                    if vals[x] != vals[x - 1]:
                        best = min(best, vals[x] - vals[x - 1])

                ans[i][j] = 0 if best == float('inf') else best

        return ans