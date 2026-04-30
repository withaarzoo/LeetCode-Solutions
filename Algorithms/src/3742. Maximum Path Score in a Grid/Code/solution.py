from typing import List

class Solution:
    def maxPathScore(self, grid: List[List[int]], k: int) -> int:
        m = len(grid)                         # Number of rows in the grid.
        n = len(grid[0])                      # Number of columns in the grid.
        NEG = -10**9                          # Sentinel for impossible states.

        # prev[j][c] = best score at column j in the previous row with exact cost c.
        prev = [[NEG] * (k + 1) for _ in range(n)]

        for i in range(m):
            # Rebuild the current row from scratch so old values do not interfere.
            curr = [[NEG] * (k + 1) for _ in range(n)]

            for j in range(n):
                gain = grid[i][j]             # Score gained by taking this cell.
                need = 1 if gain > 0 else 0   # Cost spent by this cell: 0 for 0, 1 for 1/2.

                # A path to (i, j) cannot spend more than i + j budget points.
                limit = min(k, i + j)

                # The first cell is the base case.
                if i == 0 and j == 0:
                    curr[0][0] = 0            # Start with zero score and zero cost.
                    continue

                for c in range(need, limit + 1):
                    best = NEG

                    # Take the path from above, then pay for this cell.
                    if i > 0 and prev[j][c - need] != NEG:
                        best = max(best, prev[j][c - need] + gain)

                    # Take the path from the left, then pay for this cell.
                    if j > 0 and curr[j - 1][c - need] != NEG:
                        best = max(best, curr[j - 1][c - need] + gain)

                    curr[j][c] = best         # Save the best exact-cost result for this cell.

            prev = curr                       # Move the current row into prev for the next round.

        ans = max(prev[n - 1])               # Check every allowed cost at the finish cell.
        return -1 if ans < 0 else ans        # If no valid path exists, return -1.