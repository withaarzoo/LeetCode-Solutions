from typing import List

class Solution:
    def maximumScore(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if n == 1:
            return 0

        # pref[c][k] = sum of first k cells in column c
        pref = [[0] * (n + 1) for _ in range(n)]
        for c in range(n):
            s = 0
            for r in range(n):
                s += grid[r][c]
                pref[c][r + 1] = s

        NEG = -10**30

        # dp[a][b] = best score after processing up to current column,
        # with previous height = a and current height = b.
        dp = [[NEG] * (n + 1) for _ in range(n + 1)]

        # Initialize using the first column.
        for a in range(n + 1):
            for b in range(n + 1):
                dp[a][b] = max(0, pref[0][b] - pref[0][a])

        for col in range(1, n):
            ndp = [[NEG] * (n + 1) for _ in range(n + 1)]

            for mid in range(n + 1):
                # q[x] = gain of current column if the tallest neighbor height is x
                q = [max(0, pref[col][x] - pref[col][mid]) for x in range(n + 1)]

                # prefixBest[c] = max dp[a][mid] for a <= c
                prefixBest = [NEG] * (n + 1)
                prefixBest[0] = dp[0][mid]
                for a in range(1, n + 1):
                    prefixBest[a] = max(prefixBest[a - 1], dp[a][mid])

                # suffixBest[c] = max(dp[a][mid] + q[a]) for a >= c
                suffixBest = [NEG] * (n + 2)
                suffixBest[n] = dp[n][mid] + q[n]
                for a in range(n - 1, -1, -1):
                    suffixBest[a] = max(suffixBest[a + 1], dp[a][mid] + q[a])

                # For the last real column, the next height is fixed to 0.
                limit = 0 if col == n - 1 else n

                for nxt in range(limit + 1):
                    best = NEG

                    if prefixBest[nxt] != NEG:
                        best = max(best, prefixBest[nxt] + q[nxt])
                    if suffixBest[nxt + 1] != NEG:
                        best = max(best, suffixBest[nxt + 1])

                    ndp[mid][nxt] = max(ndp[mid][nxt], best)

            dp = ndp

        ans = 0
        for row in dp:
            ans = max(ans, max(row))
        return ans