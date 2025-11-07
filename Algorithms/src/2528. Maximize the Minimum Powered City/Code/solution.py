from typing import List

class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)

        # 1) Build base power with a difference array.
        diff = [0] * (n + 1)
        for i, v in enumerate(stations):
            L = max(0, i - r)
            R = min(n, i + r + 1)
            diff[L] += v
            diff[R] -= v

        base = [0] * n
        run = 0
        for i in range(n):
            run += diff[i]
            base[i] = run

        # 2) Binary search the target T.
        lo, hi = 0, sum(stations) + k
        ans = 0

        def ok(T: int) -> bool:
            added = [0] * (n + 1)
            extra = 0
            used = 0
            for i in range(n):
                extra += added[i]
                curr = base[i] + extra
                if curr < T:
                    need = T - curr
                    used += need
                    if used > k:
                        return False
                    extra += need
                    end = min(n, i + 2 * r + 1)
                    added[end] -= need
            return True

        while lo <= hi:
            mid = (lo + hi) // 2
            if ok(mid):
                ans = mid
                lo = mid + 1
            else:
                hi = mid - 1
        return ans
