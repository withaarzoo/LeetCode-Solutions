class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        c1, c2, cb = cost1, cost2, costBoth
        n1, n2 = need1, need2
        pairs = min(n1, n2)
        cost_pairs = pairs * min(cb, c1 + c2)
        rem1 = n1 - pairs
        rem2 = n2 - pairs
        cost_rem = rem1 * min(c1, cb) + rem2 * min(c2, cb)
        return cost_pairs + cost_rem