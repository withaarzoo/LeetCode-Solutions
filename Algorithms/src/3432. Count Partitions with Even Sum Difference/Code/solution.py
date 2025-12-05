from typing import List

class Solution:
    def countPartitions(self, nums: List[int]) -> int:
        # Compute the total sum of the array
        total = sum(nums)
        
        # If total sum is odd, no valid partition
        if total % 2 == 1:
            return 0
        
        # If total is even, every position between elements is a valid partition
        n = len(nums)
        return n - 1
