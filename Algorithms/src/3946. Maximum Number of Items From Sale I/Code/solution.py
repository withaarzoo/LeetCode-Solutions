class Solution:
    def maximumSaleItems(self, items: List[List[int]], budget: int) -> int:
        mult = [0] * 1501
        for it in items:
            fac = it[0]
            for f in range(1, 1501):
                if fac % f == 0:
                    mult[f] += 1
        dp = [0] * (budget + 1)
        for it in items:
            f, p = it[0], it[1]
            ndp = dp[:]
            for b in range(budget, p - 1, -1):
                ndp[b] = max(ndp[b], dp[b - p] + mult[f])
            for b in range(p, budget + 1):
                ndp[b] = max(ndp[b], ndp[b - p] + 1)
            dp = ndp
        return max(dp)