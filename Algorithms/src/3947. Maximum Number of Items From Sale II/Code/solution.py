class Solution:
    def maximumSaleItems(self, items: List[List[int]], budget: int) -> int:
        n = len(items)
        cnt = [0] * (n + 1)
        for f, p in items:
            cnt[f] += 1
        mult = [0] * (n + 1)
        for f in range(1, n + 1):
            for m in range(f, n + 1, f):
                mult[f] += cnt[m]
        min_p = min(p for f, p in items)
        cand = []
        for f, p in items:
            g = mult[f] - 1
            if g > 0:
                cand.append((p, g))
        cand.sort()
        spent = 0
        total = 0
        B = budget
        for p, g in cand:
            if p > 2 * min_p:
                break
            can = min(g, (B - spent) // p)
            total += can * 2
            spent += can * p
        total += (B - spent) // min_p
        return total