from typing import List

class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        # Total sum of all battery capacities
        total = sum(batteries)

        # Maximum time per computer cannot exceed total / n
        low, high = 0, total // n

        # Binary search on possible running time
        while low < high:
            mid = (low + high + 1) // 2  # upper mid

            usable = 0
            for b in batteries:
                # Each battery contributes at most mid minutes
                usable += min(b, mid)
                if usable >= mid * n:
                    break

            if usable >= mid * n:
                # mid minutes is possible
                low = mid
            else:
                # mid minutes is not possible
                high = mid - 1

        return low
