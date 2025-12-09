from collections import Counter
from typing import List

class Solution:
    def specialTriplets(self, nums: List[int]) -> int:
        MOD = 10**9 + 7

        right = Counter(nums)  # counts for the right side
        left = Counter()       # counts for the left side

        ans = 0

        for x in nums:
            # x is now used as the middle value, remove from right
            right[x] -= 1

            target = x * 2  # value 2x

            cnt_left = left[target]
            cnt_right = right[target]

            ans = (ans + (cnt_left * cnt_right) % MOD) % MOD

            # move x to left side
            left[x] += 1

        return ans % MOD
