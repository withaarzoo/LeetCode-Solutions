import math
from typing import List

class Solution:
    def minimizedMaximum(self, n: int, quantities: List[int]) -> int:
        def canDistribute(maxProducts):
            storesNeeded = 0
            for quantity in quantities:
                storesNeeded += math.ceil(quantity / maxProducts)
                if storesNeeded > n:
                    return False
            return storesNeeded <= n

        low, high = 1, max(quantities)
        answer = high

        while low <= high:
            mid = (low + high) // 2
            if canDistribute(mid):
                answer = mid
                high = mid - 1
            else:
                low = mid + 1

        return answer