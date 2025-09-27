from typing import List

class Solution:
    def largestTriangleArea(self, points: List[List[int]]) -> float:
        n = len(points)
        max_area = 0.0
        # iterate over all combinations of three distinct points
        for i in range(n - 2):
            for j in range(i + 1, n - 1):
                for k in range(j + 1, n):
                    x1, y1 = points[i]
                    x2, y2 = points[j]
                    x3, y3 = points[k]
                    # shoelace/cross product formula gives doubled area
                    doubled = abs(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2))
                    area = doubled * 0.5
                    if area > max_area:
                        max_area = area
        return max_area
