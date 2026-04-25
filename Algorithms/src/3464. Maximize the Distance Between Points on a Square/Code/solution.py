from bisect import bisect_left
from typing import List

class Solution:
    def maxDistance(self, side: int, points: List[List[int]], k: int) -> int:
        pts = []
        for x, y in points:
            if x == 0:
                pos = y
            elif y == side:
                pos = side + x
            elif x == side:
                pos = 3 * side - y
            else:
                pos = 4 * side - x
            pts.append((pos, x, y))

        pts.sort()
        n = len(pts)

        def get_offset(x: int, y: int, d: int) -> int:
            if x == 0:  # left
                if d <= 2 * side - y:
                    return d
                if d <= side + y:
                    return 2 * side + d - 2 * y
                return -1
            elif y == side:  # top
                if d <= 2 * side - x:
                    return d
                if d <= side + x:
                    return 2 * side + d - 2 * x
                return -1
            elif x == side:  # right
                if d <= side + y:
                    return d
                if d <= 2 * side - y:
                    return d + 2 * y
                return -1
            else:  # bottom
                if d <= side + x:
                    return d
                if d <= 2 * side - x:
                    return d + 2 * x
                return -1

        def feasible(d: int) -> bool:
            pos = [p[0] for p in pts]
            pos3 = pos + [p + 4 * side for p in pos] + [p + 8 * side for p in pos]

            nxt = [-1] * (2 * n)

            for i in range(2 * n):
                _, x, y = pts[i % n]
                off = get_offset(x, y, d)
                if off < 0:
                    continue

                target = pos3[i] + off
                hi = min(i + n, 3 * n)
                j = bisect_left(pos3, target, i + 1, hi)
                if j < hi:
                    nxt[i] = j

            for start in range(n):
                cur = start
                cnt = 1

                while cnt < k:
                    cur = nxt[cur]
                    if cur == -1 or cur >= start + n:
                        break
                    cnt += 1

                if cnt >= k:
                    x1, y1 = pts[start][1], pts[start][2]
                    x2, y2 = pts[cur % n][1], pts[cur % n][2]
                    if abs(x1 - x2) + abs(y1 - y2) >= d:
                        return True

            return False

        lo, hi = 0, 2 * side
        while lo < hi:
            mid = (lo + hi + 1) // 2
            if feasible(mid):
                lo = mid
            else:
                hi = mid - 1

        return lo