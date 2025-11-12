from math import gcd
from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)

        # 1) If array already has ones, convert the rest.
        ones = sum(1 for x in nums if x == 1)
        if ones > 0:
            return n - ones

        # 2) If global gcd > 1, we can never reach 1.
        g = 0
        for x in nums:
            g = gcd(g, x)
        if g > 1:
            return -1

        # 3) Find shortest subarray with gcd == 1.
        best = 10**9
        for i in range(n):
            cur = 0
            for j in range(i, n):
                cur = gcd(cur, nums[j])
                if cur == 1:
                    best = min(best, j - i + 1)
                    break

        # 4) Make first 1 (best-1 ops) + spread to all (n-1 ops).
        return (best - 1) + (n - 1)
