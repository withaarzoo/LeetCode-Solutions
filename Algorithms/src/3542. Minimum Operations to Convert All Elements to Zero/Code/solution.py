from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        stk = []  # non-decreasing stack
        ans = 0
        for x in nums:
            while stk and stk[-1] > x:
                stk.pop()
            if x == 0:
                continue
            if not stk or stk[-1] < x:
                ans += 1
                stk.append(x)
        return ans
