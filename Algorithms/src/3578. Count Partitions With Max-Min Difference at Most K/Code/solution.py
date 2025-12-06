from collections import deque
from typing import List

class Solution:
    def countPartitions(self, nums: List[int], k: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)

        # dp[i]: ways to partition nums[0..i-1]
        dp = [0] * (n + 1)
        pref = [0] * (n + 1)

        dp[0] = 1
        pref[0] = 1

        maxdq = deque()  # indices, nums decreasing
        mindq = deque()  # indices, nums increasing
        l = 0

        for r in range(n):
            x = nums[r]

            # maintain decreasing deque for max
            while maxdq and nums[maxdq[-1]] <= x:
                maxdq.pop()
            maxdq.append(r)

            # maintain increasing deque for min
            while mindq and nums[mindq[-1]] >= x:
                mindq.pop()
            mindq.append(r)

            # move l until window [l..r] is valid
            while maxdq and mindq and nums[maxdq[0]] - nums[mindq[0]] > k:
                if maxdq[0] == l:
                    maxdq.popleft()
                if mindq[0] == l:
                    mindq.popleft()
                l += 1

            L = l
            i = r + 1

            ways = pref[i - 1]
            if L > 0:
                ways -= pref[L - 1]
            ways %= MOD

            dp[i] = ways
            pref[i] = (pref[i - 1] + dp[i]) % MOD

        return dp[n]
