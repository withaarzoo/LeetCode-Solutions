from typing import List
from functools import lru_cache

class Solution:
    def stoneGameII(self, piles: List[int]) -> int:
        n = len(piles)
        suffix = [0] * (n + 1)

        for i in range(n - 1, -1, -1):
            suffix[i] = suffix[i + 1] + piles[i]

        @lru_cache(None)
        def solve(i, M):
            if i == n:
                return 0

            best = 0

            for X in range(1, min(2 * M, n - i) + 1):
                next_M = max(M, X)
                current = suffix[i] - solve(i + X, next_M)
                best = max(best, current)

            return best

        return solve(0, 1)