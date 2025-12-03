from typing import List
from collections import defaultdict
from math import gcd

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        n = len(points)
        SHIFT = 3000
        
        def encode_pair(a: int, b: int) -> int:
            return ((a + SHIFT) << 13) ^ (b + SHIFT)
        
        # slope_key -> { line_id: count }
        by_slope = defaultdict(lambda: defaultdict(int))
        # vector_key -> { line_id: count }
        by_vector = defaultdict(lambda: defaultdict(int))
        
        for i in range(n):
            x1, y1 = points[i]
            for j in range(i + 1, n):
                x2, y2 = points[j]
                dx = x2 - x1
                dy = y2 - y1
                
                if dx < 0 or (dx == 0 and dy < 0):
                    dx = -dx
                    dy = -dy
                
                g = gcd(dx, dy)
                ux = dx // g
                uy = dy // g
                
                line_id = ux * y1 - uy * x1
                
                slope_key = encode_pair(ux, uy)
                vector_key = encode_pair(dx, dy)
                
                by_slope[slope_key][line_id] += 1
                by_vector[vector_key][line_id] += 1
        
        def count_pairs(mp) -> int:
            ans = 0
            for inner in mp.values():
                counts = inner.values()
                s = sum(counts)
                sum_sq = sum(c * c for c in counts)
                ans += (s * s - sum_sq) // 2
            return ans
        
        with_parallel = count_pairs(by_slope)
        parallelogram_twice = count_pairs(by_vector)
        
        return with_parallel - parallelogram_twice // 2
