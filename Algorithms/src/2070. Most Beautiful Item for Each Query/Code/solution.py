from bisect import bisect_right
from typing import List

class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort()
        price_beauty = []
        max_beauty = 0

        for price, beauty in items:
            max_beauty = max(max_beauty, beauty)
            price_beauty.append((price, max_beauty))

        result = []
        for query in queries:
            idx = bisect_right(price_beauty, (query, float('inf'))) - 1
            result.append(price_beauty[idx][1] if idx >= 0 else 0)
        
        return result