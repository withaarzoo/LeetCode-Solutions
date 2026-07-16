from math import gcd

class Solution:
    def gcdSum(self, nums: list[int]) -> int:
        n = len(nums)

        # Store prefix gcd values
        prefix_gcd = []

        # Running prefix maximum
        prefix_max = 0

        # Build prefix_gcd
        for x in nums:
            prefix_max = max(prefix_max, x)
            prefix_gcd.append(gcd(x, prefix_max))

        # Sort the array
        prefix_gcd.sort()

        ans = 0

        # Pair smallest with largest
        left = 0
        right = n - 1

        while left < right:
            ans += gcd(prefix_gcd[left], prefix_gcd[right])
            left += 1
            right -= 1

        return ans