from typing import List
from collections import defaultdict

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        MOD = 10**9 + 7
        INV2 = (MOD + 1) // 2  # modular inverse of 2

        # 1. Count points per y-coordinate
        freq = defaultdict(int)
        for x, y in points:
            freq[y] += 1

        sumF = 0  # S
        sumF2 = 0 # SQ

        # 2. For each y, compute C(c,2) and accumulate
        for c in freq.values():
            if c >= 2:
                f = c * (c - 1) // 2  # C(c,2)
                f %= MOD
                sumF = (sumF + f) % MOD
                sumF2 = (sumF2 + f * f) % MOD

        # 3. ((S^2 - SQ) / 2) mod MOD
        ans = (sumF * sumF) % MOD
        ans = (ans - sumF2 + MOD) % MOD
        ans = ans * INV2 % MOD

        return ans
