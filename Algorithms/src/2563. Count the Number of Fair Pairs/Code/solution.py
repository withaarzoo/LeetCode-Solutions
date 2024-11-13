from bisect import bisect_left, bisect_right
from typing import List

class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        nums.sort()
        count = 0
        n = len(nums)
        
        for i in range(n - 1):
            min_val = lower - nums[i]
            max_val = upper - nums[i]
            
            # Using bisect to find the range
            start = bisect_left(nums, min_val, i + 1)
            end = bisect_right(nums, max_val, i + 1)
            
            count += (end - start)
        
        return count