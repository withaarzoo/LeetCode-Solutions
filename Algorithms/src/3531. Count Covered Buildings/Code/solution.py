from typing import List
import bisect

class Solution:
    def countCoveredBuildings(self, n: int, buildings: List[List[int]]) -> int:
        # maps: row x -> list of y, col y -> list of x
        row = {}
        col = {}
        for x, y in buildings:
            row.setdefault(x, []).append(y)
            col.setdefault(y, []).append(x)

        # sort each list
        for ys in row.values():
            ys.sort()
        for xs in col.values():
            xs.sort()

        ans = 0
        for x, y in buildings:
            ys = row[x]
            xs = col[y]
            posY = bisect.bisect_left(ys, y)
            posX = bisect.bisect_left(xs, x)
            insideRow = (posY > 0 and posY < len(ys) - 1)
            insideCol = (posX > 0 and posX < len(xs) - 1)
            if insideRow and insideCol:
                ans += 1
        return ans
