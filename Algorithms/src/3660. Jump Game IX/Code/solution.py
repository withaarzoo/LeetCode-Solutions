from typing import List

class Solution:
    def maxValue(self, nums: List[int]) -> List[int]:
        n = len(nums)

        # suffixMin[i] = smallest value in nums[i...n-1]
        # I keep one extra cell so the last segment can close cleanly.
        suffixMin = [0] * (n + 1)
        suffixMin[n] = float("inf")
        for i in range(n - 1, -1, -1):
            suffixMin[i] = min(nums[i], suffixMin[i + 1])

        ans = [0] * n
        l = 0

        # I walk through the array and build one connected segment at a time.
        while l < n:
            r = l
            component_max = nums[l]

            # I extend the segment while some inversion still crosses the next cut.
            while r + 1 < n and component_max > suffixMin[r + 1]:
                r += 1
                component_max = max(component_max, nums[r])

            # Every index in this segment can reach this segment maximum.
            for i in range(l, r + 1):
                ans[i] = component_max

            # Continue with the next segment.
            l = r + 1

        return ans