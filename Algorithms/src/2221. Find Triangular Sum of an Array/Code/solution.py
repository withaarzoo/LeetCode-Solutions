from typing import List

class Solution:
    def triangularSum(self, nums: List[int]) -> int:
        n = len(nums)
        # shrink the effective length from n to 1
        for length in range(n, 1, -1):
            # compute new values in-place into nums[0..length-2]
            for i in range(length - 1):
                nums[i] = (nums[i] + nums[i + 1]) % 10
        return nums[0]
