from typing import List

class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        n = len(nums)
        # Build intervals [num-k, num+k]
        intervals = [(x - k, x + k) for x in nums]
        # Sort by right endpoint then left
        intervals.sort(key=lambda t: (t[1], t[0]))
        last_assigned = -10**30  # very small
        ans = 0
        for l, r in intervals:
            assigned = max(l, last_assigned + 1)
            if assigned <= r:
                ans += 1
                last_assigned = assigned
        return ans
